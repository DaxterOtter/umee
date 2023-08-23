package suite

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/umee-network/umee/v6/app"

	"github.com/cosmos/cosmos-sdk/server"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	bech32ibctypes "github.com/osmosis-labs/bech32-ibc/x/bech32ibc/types"
	"github.com/spf13/viper"
	tmconfig "github.com/tendermint/tendermint/config"
	tmjson "github.com/tendermint/tendermint/libs/json"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	appparams "github.com/umee-network/umee/v6/app/params"
	"github.com/umee-network/umee/v6/x/leverage/fixtures"
	leveragetypes "github.com/umee-network/umee/v6/x/leverage/types"
	oracletypes "github.com/umee-network/umee/v6/x/oracle/types"
	"github.com/umee-network/umee/v6/x/uibc"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/suite"
	"github.com/umee-network/umee/v6/client"
)

type SimpleSuite struct {
	suite.Suite

	Chain        *SimpleChain
	DkrPool      *dockertest.Pool
	DkrNet       *dockertest.Network
	ValResources []*dockertest.Resource
	Umee         client.Client
	Cdc          codec.Codec
}

func NewSimpleSuite(numValidators int) *SimpleSuite {
	var err error
	var s *SimpleSuite
	s.T().Log("setting up simple test suite...")

	// codec
	s.Cdc = EncodingConfig.Codec

	s.Chain, err = NewSimpleChain()
	s.Require().NoError(err)

	s.T().Logf("starting infrastructure; chain-id: %s; datadir: %s", s.Chain.ID, s.Chain.DataDir)

	s.DkrPool, err = dockertest.NewPool("")
	s.Require().NoError(err)

	s.DkrNet, err = s.DkrPool.CreateNetwork(fmt.Sprintf("%s-testnet", s.Chain.ID))
	s.Require().NoError(err)

	s.initNodes(numValidators) // init validator nodes
	s.initGenesis()
	s.initValidatorConfigs()
	s.runValidators()
	s.initUmeeClient()

	return s
}

func (s *SimpleSuite) TearDownSuite() {
	if str := os.Getenv("UMEE_E2E_SKIP_CLEANUP"); len(str) > 0 {
		skipCleanup, err := strconv.ParseBool(str)
		s.Require().NoError(err)

		if skipCleanup {
			return
		}
	}

	s.T().Log("tearing down simple test suite...")

	for _, vc := range s.ValResources {
		s.Require().NoError(s.DkrPool.Purge(vc))
	}

	s.Require().NoError(s.DkrPool.RemoveNetwork(s.DkrNet))
}

func (s *SimpleSuite) initNodes(numValidators int) {
	s.Require().NoError(s.Chain.CreateAndInitValidators(s.Cdc, numValidators))

	// initialize a genesis file for the first validator
	val0ConfigDir := s.Chain.Validators[0].ConfigDir()
	for _, val := range s.Chain.Validators {
		valAddr, err := val.KeyInfo.GetAddress()
		s.Require().NoError(err)
		s.Require().NoError(
			AddGenesisAccount(s.Cdc, val0ConfigDir, "", InitBalanceStr, valAddr),
		)
	}

	// copy the genesis file to the remaining validators
	for _, val := range s.Chain.Validators[1:] {
		_, err := CopyFile(
			filepath.Join(val0ConfigDir, "config", "genesis.json"),
			filepath.Join(val.ConfigDir(), "config", "genesis.json"),
		)
		s.Require().NoError(err)
	}
}

func (s *SimpleSuite) initGenesis() {
	serverCtx := server.NewDefaultContext()
	config := serverCtx.Config

	config.SetRoot(s.Chain.Validators[0].ConfigDir())
	config.Moniker = s.Chain.Validators[0].Moniker

	genFilePath := config.GenesisFile()
	s.T().Log("starting e2e infrastructure; validator_0 config:", genFilePath)
	appGenState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFilePath)
	s.Require().NoError(err)

	var bech32GenState bech32ibctypes.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[bech32ibctypes.ModuleName], &bech32GenState))

	// bech32
	bech32GenState.NativeHRP = sdk.GetConfig().GetBech32AccountAddrPrefix()

	bz, err := s.Cdc.MarshalJSON(&bech32GenState)
	s.Require().NoError(err)
	appGenState[bech32ibctypes.ModuleName] = bz

	// Leverage
	var leverageGenState leveragetypes.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[leveragetypes.ModuleName], &leverageGenState))

	leverageGenState.Registry = append(leverageGenState.Registry,
		fixtures.Token(appparams.BondDenom, appparams.DisplayDenom, 6),
		fixtures.Token(ATOMBaseDenom, ATOM, uint32(ATOMExponent)),
	)

	bz, err = s.Cdc.MarshalJSON(&leverageGenState)
	s.Require().NoError(err)
	appGenState[leveragetypes.ModuleName] = bz

	// Oracle
	var oracleGenState oracletypes.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[oracletypes.ModuleName], &oracleGenState))

	oracleGenState.Params.HistoricStampPeriod = 5
	oracleGenState.Params.MaximumPriceStamps = 4
	oracleGenState.Params.MedianStampPeriod = 20
	oracleGenState.Params.MaximumMedianStamps = 2

	oracleGenState.Params.AcceptList = append(oracleGenState.Params.AcceptList, oracletypes.Denom{
		BaseDenom:   ATOMBaseDenom,
		SymbolDenom: ATOM,
		Exponent:    uint32(ATOMExponent),
	})

	bz, err = s.Cdc.MarshalJSON(&oracleGenState)
	s.Require().NoError(err)
	appGenState[oracletypes.ModuleName] = bz

	// Gov
	var govGenState govtypesv1.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[govtypes.ModuleName], &govGenState))

	votingPeriod := 5 * time.Second
	govGenState.VotingParams.VotingPeriod = &votingPeriod
	govGenState.DepositParams.MinDeposit = sdk.NewCoins(sdk.NewCoin(appparams.BondDenom, sdk.NewInt(100)))

	bz, err = s.Cdc.MarshalJSON(&govGenState)
	s.Require().NoError(err)
	appGenState[govtypes.ModuleName] = bz

	// Bank
	var bankGenState banktypes.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[banktypes.ModuleName], &bankGenState))

	bankGenState.DenomMetadata = append(bankGenState.DenomMetadata, banktypes.Metadata{
		Description: "An example stable token",
		Display:     PhotonDenom,
		Base:        PhotonDenom,
		Symbol:      PhotonDenom,
		Name:        PhotonDenom,
		DenomUnits: []*banktypes.DenomUnit{
			{
				Denom:    PhotonDenom,
				Exponent: 0,
			},
		},
	})

	bz, err = s.Cdc.MarshalJSON(&bankGenState)
	s.Require().NoError(err)
	appGenState[banktypes.ModuleName] = bz

	// uibc (ibc quota)
	var uibcGenState uibc.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[uibc.ModuleName], &uibcGenState))

	// 100$ for each token
	uibcGenState.Params.TokenQuota = sdk.NewDec(100)
	// 120$ for all tokens on quota duration
	uibcGenState.Params.TotalQuota = sdk.NewDec(120)
	// quotas will reset every 300 seconds
	uibcGenState.Params.QuotaDuration = time.Second * 300

	bz, err = s.Cdc.MarshalJSON(&uibcGenState)
	s.Require().NoError(err)
	appGenState[uibc.ModuleName] = bz

	var genUtilGenState genutiltypes.GenesisState
	s.Require().NoError(s.Cdc.UnmarshalJSON(appGenState[genutiltypes.ModuleName], &genUtilGenState))

	// generate genesis txs
	genTxs := make([]json.RawMessage, len(s.Chain.Validators))
	for i, val := range s.Chain.Validators {
		var createValmsg sdk.Msg
		if i == 2 {
			createValmsg, err = val.BuildCreateValidatorMsg(stakeAmountCoin2)
		} else {
			createValmsg, err = val.BuildCreateValidatorMsg(stakeAmountCoin)
		}
		s.Require().NoError(err)

		signedTx, err := val.SignMsg(s.Cdc, createValmsg)
		s.Require().NoError(err)

		txRaw, err := s.Cdc.MarshalJSON(signedTx)
		s.Require().NoError(err)

		genTxs[i] = txRaw
	}

	genUtilGenState.GenTxs = genTxs

	bz, err = s.Cdc.MarshalJSON(&genUtilGenState)
	s.Require().NoError(err)
	appGenState[genutiltypes.ModuleName] = bz

	bz, err = json.MarshalIndent(appGenState, "", "  ")
	s.Require().NoError(err)

	genDoc.AppState = bz

	bz, err = tmjson.MarshalIndent(genDoc, "", "  ")
	s.Require().NoError(err)

	// write the updated genesis file to each validator
	for _, val := range s.Chain.Validators {
		writeFile(filepath.Join(val.ConfigDir(), "config", "genesis.json"), bz)
	}
}

func (s *SimpleSuite) initValidatorConfigs() {
	for i, val := range s.Chain.Validators {
		tmCfgPath := filepath.Join(val.ConfigDir(), "config", "config.toml")

		vpr := viper.New()
		vpr.SetConfigFile(tmCfgPath)
		s.Require().NoError(vpr.ReadInConfig())

		valConfig := tmconfig.DefaultConfig()
		valConfig.Consensus.SkipTimeoutCommit = true
		s.Require().NoError(vpr.Unmarshal(valConfig))

		valConfig.P2P.ListenAddress = "tcp://0.0.0.0:26656"
		valConfig.P2P.AddrBookStrict = false
		valConfig.P2P.ExternalAddress = fmt.Sprintf("%s:%d", val.InstanceName(), 26656)
		valConfig.RPC.ListenAddress = "tcp://0.0.0.0:26657"
		valConfig.StateSync.Enable = false
		valConfig.LogLevel = "info"

		var peers []string

		for j := 0; j < len(s.Chain.Validators); j++ {
			if i == j {
				continue
			}

			peer := s.Chain.Validators[j]
			peerID := fmt.Sprintf("%s@%s%d:26656", peer.NodeKey.ID(), peer.Moniker, j)
			peers = append(peers, peerID)
		}

		valConfig.P2P.PersistentPeers = strings.Join(peers, ",")

		tmconfig.WriteConfigFile(tmCfgPath, valConfig)

		// set application configuration
		appCfgPath := filepath.Join(val.ConfigDir(), "config", "app.toml")

		appConfig := srvconfig.DefaultConfig()
		appConfig.API.Enable = true
		appConfig.MinGasPrices = minGasPrice
		appConfig.Pruning = "nothing"

		srvconfig.WriteConfigFile(appCfgPath, appConfig)
	}
}

func (s *SimpleSuite) runValidators() {
	s.T().Log("starting Umee validator containers...")

	s.ValResources = make([]*dockertest.Resource, len(s.Chain.Validators))
	for i, val := range s.Chain.Validators {
		runOpts := &dockertest.RunOptions{
			Name:      val.InstanceName(),
			NetworkID: s.DkrNet.Network.ID,
			Mounts: []string{
				fmt.Sprintf("%s/:/root/.umee", val.ConfigDir()),
			},
			Repository: "umee-network/umeed-e2e",
		}

		// expose the first validator for debugging and communication
		if val.Index == 0 {
			runOpts.PortBindings = map[docker.Port][]docker.PortBinding{
				"1317/tcp":  {{HostIP: "", HostPort: "1317"}},
				"6060/tcp":  {{HostIP: "", HostPort: "6060"}},
				"6061/tcp":  {{HostIP: "", HostPort: "6061"}},
				"6062/tcp":  {{HostIP: "", HostPort: "6062"}},
				"6063/tcp":  {{HostIP: "", HostPort: "6063"}},
				"6064/tcp":  {{HostIP: "", HostPort: "6064"}},
				"6065/tcp":  {{HostIP: "", HostPort: "6065"}},
				"9090/tcp":  {{HostIP: "", HostPort: "9090"}},
				"26656/tcp": {{HostIP: "", HostPort: "26656"}},
				"26657/tcp": {{HostIP: "", HostPort: "26657"}},
			}
		}

		resource, err := s.DkrPool.RunWithOptions(runOpts, noRestart)
		s.Require().NoError(err)

		s.ValResources[i] = resource
		s.T().Logf("started Umee validator container: %s", resource.Container.ID)
	}

	rpcClient, err := rpchttp.New("tcp://localhost:26657", "/websocket")
	s.Require().NoError(err)

	s.Require().Eventually(
		func() bool {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			status, err := rpcClient.Status(ctx)
			if err != nil {
				return false
			}

			// let the node produce a few blocks
			if status.SyncInfo.CatchingUp || status.SyncInfo.LatestBlockHeight < 3 {
				return false
			}

			return true
		},
		5*time.Minute,
		time.Second,
		"umee node failed to produce blocks",
	)
}

func (s *SimpleSuite) initUmeeClient() {
	var err error
	mnemonics := make(map[string]string)
	for index, v := range s.Chain.Validators {
		mnemonics[fmt.Sprintf("val%d", index)] = v.Mnemonic
	}
	ecfg := app.MakeEncodingConfig()
	s.Umee, err = client.NewClient(
		s.Chain.DataDir,
		s.Chain.ID,
		"tcp://localhost:26657",
		"tcp://localhost:9090",
		mnemonics,
		1,
		ecfg,
	)
	s.Require().NoError(err)
}

func noRestart(config *docker.HostConfig) {
	// in this case we don't want the nodes to restart on failure
	config.RestartPolicy = docker.RestartPolicy{
		Name: "no",
	}
}
