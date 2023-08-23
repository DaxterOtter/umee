package setup

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"

	ssuite "github.com/umee-network/umee/v6/tests/suite"
)

type E2ETestSuite struct {
	*ssuite.SimpleSuite

	tmpDirs             []string
	Chain               *e2eChain
	gaiaRPC             *rpchttp.HTTP
	GaiaResource        *dockertest.Resource
	HermesResource      *dockertest.Resource
	priceFeederResource *dockertest.Resource
}

func (s *E2ETestSuite) SetupSuite() {
	s.SimpleSuite = ssuite.NewSimpleSuite(3)
	s.T().Log("setting up e2e integration test suite...")

	s.Chain = &e2eChain{
		SimpleChain: s.SimpleSuite.Chain,
	}
	s.T().Logf("starting e2e infrastructure; e2eChain-id: %s; datadir: %s", s.Chain.ID, s.Chain.DataDir)

	s.Require().NoError(s.Chain.createAndInitGaiaValidator(s.Cdc))
	s.runPriceFeeder()
	s.runGaiaNetwork()
	time.Sleep(3 * time.Second) // wait for gaia to start
	s.runIBCRelayer()
}

func (s *E2ETestSuite) TearDownSuite() {
	if str := os.Getenv("UMEE_E2E_SKIP_CLEANUP"); len(str) > 0 {
		skipCleanup, err := strconv.ParseBool(str)
		s.Require().NoError(err)

		if skipCleanup {
			return
		}
	}

	s.T().Log("tearing down e2e integration test suite...")

	s.Require().NoError(s.DkrPool.Purge(s.GaiaResource))
	s.Require().NoError(s.DkrPool.Purge(s.HermesResource))
	s.Require().NoError(s.DkrPool.Purge(s.priceFeederResource))

	os.RemoveAll(s.Chain.DataDir)
	for _, td := range s.tmpDirs {
		os.RemoveAll(td)
	}
}

func (s *E2ETestSuite) runGaiaNetwork() {
	s.T().Log("starting Gaia network container...")

	tmpDir, err := os.MkdirTemp("", "umee-e2e-testnet-gaia-")
	s.Require().NoError(err)
	s.tmpDirs = append(s.tmpDirs, tmpDir)

	gaiaVal := s.Chain.GaiaValidators[0]

	gaiaCfgPath := path.Join(tmpDir, "cfg")
	s.Require().NoError(os.MkdirAll(gaiaCfgPath, 0o755))

	_, err = ssuite.CopyFile(
		filepath.Join("./scripts/", "gaia_bootstrap.sh"),
		filepath.Join(gaiaCfgPath, "gaia_bootstrap.sh"),
	)
	s.Require().NoError(err)

	s.GaiaResource, err = s.DkrPool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       gaiaVal.instanceName(),
			Repository: "ghcr.io/umee-network/gaia-e2e",
			Tag:        "latest",
			NetworkID:  s.DkrNet.Network.ID,
			Mounts: []string{
				fmt.Sprintf("%s/:/root/.gaia", tmpDir),
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"1317/tcp":  {{HostIP: "", HostPort: "1417"}},
				"9090/tcp":  {{HostIP: "", HostPort: "9190"}},
				"26656/tcp": {{HostIP: "", HostPort: "27656"}},
				"26657/tcp": {{HostIP: "", HostPort: "27657"}},
			},
			Env: []string{
				fmt.Sprintf("UMEE_E2E_GAIA_CHAIN_ID=%s", GaiaChainID),
				fmt.Sprintf("UMEE_E2E_GAIA_VAL_MNEMONIC=%s", gaiaVal.mnemonic),
			},
			Entrypoint: []string{
				"sh",
				"-c",
				"chmod +x /root/.gaia/cfg/gaia_bootstrap.sh && /root/.gaia/cfg/gaia_bootstrap.sh",
			},
		},
		noRestart,
	)
	s.Require().NoError(err)

	endpoint := fmt.Sprintf("tcp://%s", s.GaiaResource.GetHostPort("26657/tcp"))
	s.gaiaRPC, err = rpchttp.New(endpoint, "/websocket")
	s.Require().NoError(err)

	s.Require().Eventually(
		func() bool {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			status, err := s.gaiaRPC.Status(ctx)
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
		"gaia node failed to produce blocks",
	)

	s.T().Logf("started Gaia network container: %s", s.GaiaResource.Container.ID)
}

func (s *E2ETestSuite) runIBCRelayer() {
	s.T().Log("starting Hermes relayer container...")

	tmpDir, err := os.MkdirTemp("", "umee-e2e-testnet-hermes-")
	s.Require().NoError(err)
	s.tmpDirs = append(s.tmpDirs, tmpDir)

	gaiaVal := s.Chain.GaiaValidators[0]
	// umeeVal for the relayer needs to be a different account
	// than what we use for runPriceFeeder.
	umeeVal := s.Chain.Validators[1]
	hermesCfgPath := path.Join(tmpDir, "hermes")

	s.Require().NoError(os.MkdirAll(hermesCfgPath, 0o755))
	_, err = ssuite.CopyFile(
		filepath.Join("./scripts/", "hermes_bootstrap.sh"),
		filepath.Join(hermesCfgPath, "hermes_bootstrap.sh"),
	)
	s.Require().NoError(err)

	s.HermesResource, err = s.DkrPool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       "umee-gaia-relayer",
			Repository: "ghcr.io/umee-network/hermes-e2e",
			Tag:        "latest",
			NetworkID:  s.DkrNet.Network.ID,
			Mounts: []string{
				fmt.Sprintf("%s/:/home/hermes", hermesCfgPath),
			},
			ExposedPorts: []string{"3031"},
			PortBindings: map[docker.Port][]docker.PortBinding{
				"3031/tcp": {{HostIP: "", HostPort: "3031"}},
			},
			Env: []string{
				fmt.Sprintf("UMEE_E2E_GAIA_CHAIN_ID=%s", GaiaChainID),
				fmt.Sprintf("UMEE_E2E_UMEE_CHAIN_ID=%s", s.Chain.ID),
				fmt.Sprintf("UMEE_E2E_GAIA_VAL_MNEMONIC=%s", gaiaVal.mnemonic),
				fmt.Sprintf("UMEE_E2E_UMEE_VAL_MNEMONIC=%s", umeeVal.Mnemonic),
				fmt.Sprintf("UMEE_E2E_GAIA_VAL_HOST=%s", s.GaiaResource.Container.Name[1:]),
				fmt.Sprintf("UMEE_E2E_UMEE_VAL_HOST=%s", s.ValResources[1].Container.Name[1:]),
			},
			Entrypoint: []string{
				"sh",
				"-c",
				"chmod +x /home/hermes/hermes_bootstrap.sh && /home/hermes/hermes_bootstrap.sh",
			},
		},
		noRestart,
	)
	s.Require().NoError(err)

	endpoint := fmt.Sprintf("http://%s/state", s.HermesResource.GetHostPort("3031/tcp"))
	s.Require().Eventually(
		func() bool {
			resp, err := http.Get(endpoint)
			if err != nil {
				return false
			}

			defer resp.Body.Close()

			bz, err := io.ReadAll(resp.Body)
			if err != nil {
				return false
			}

			var respBody map[string]interface{}
			if err := json.Unmarshal(bz, &respBody); err != nil {
				return false
			}

			status := respBody["status"].(string)
			result := respBody["result"].(map[string]interface{})

			return status == "success" && len(result["chains"].([]interface{})) == 2
		},
		5*time.Minute,
		time.Second,
		"hermes relayer not healthy",
	)

	s.T().Logf("started Hermes relayer container: %s", s.HermesResource.Container.ID)

	// create the client, connection and channel between the Umee and Gaia chains
	s.connectIBCChains()
}

func (s *E2ETestSuite) runPriceFeeder() {
	s.T().Log("starting price-feeder container...")

	umeeVal := s.Chain.Validators[2]
	umeeValAddr, err := umeeVal.KeyInfo.GetAddress()
	s.Require().NoError(err)

	grpcEndpoint := fmt.Sprintf("tcp://%s:%s", umeeVal.InstanceName(), "9090")
	tmrpcEndpoint := fmt.Sprintf("http://%s:%s", umeeVal.InstanceName(), "26657")

	s.priceFeederResource, err = s.DkrPool.RunWithOptions(
		&dockertest.RunOptions{
			Name:       "umee-price-feeder",
			NetworkID:  s.DkrNet.Network.ID,
			Repository: PriceFeederContainerRepo,
			Mounts: []string{
				fmt.Sprintf("%s/:/root/.umee", umeeVal.ConfigDir()),
			},
			PortBindings: map[docker.Port][]docker.PortBinding{
				PriceFeederServerPort: {{HostIP: "", HostPort: "7171"}},
			},
			Env: []string{
				fmt.Sprintf("PRICE_FEEDER_PASS=%s", keyringPassphrase),
				fmt.Sprintf("ACCOUNT_ADDRESS=%s", umeeValAddr),
				fmt.Sprintf("ACCOUNT_VALIDATOR=%s", sdk.ValAddress(umeeValAddr)),
				fmt.Sprintf("KEYRING_DIR=%s", "/root/.umee"),
				fmt.Sprintf("ACCOUNT_CHAIN_ID=%s", s.Chain.ID),
				fmt.Sprintf("RPC_GRPC_ENDPOINT=%s", grpcEndpoint),
				fmt.Sprintf("RPC_TMRPC_ENDPOINT=%s", tmrpcEndpoint),
			},
			Cmd: []string{
				"--skip-provider-check",
				"--log-level=debug",
			},
		},
		noRestart,
	)
	s.Require().NoError(err)

	endpoint := fmt.Sprintf("http://%s/api/v1/prices", s.priceFeederResource.GetHostPort(PriceFeederServerPort))

	checkHealth := func() bool {
		resp, err := http.Get(endpoint)
		if err != nil {
			s.T().Log("Price feeder endpoint not available", err)
			return false
		}

		defer resp.Body.Close()

		bz, err := io.ReadAll(resp.Body)
		if err != nil {
			s.T().Log("Can't get price feeder response", err)
			return false
		}

		var respBody map[string]interface{}
		if err := json.Unmarshal(bz, &respBody); err != nil {
			s.T().Log("Can't unmarshal price feed", err)
			return false
		}

		prices, ok := respBody["prices"].(map[string]interface{})
		if !ok {
			s.T().Log("price feeder: no prices")
			return false
		}

		return len(prices) > 0
	}

	isHealthy := false
	for i := 0; i < PriceFeederMaxStartupTime; i++ {
		isHealthy = checkHealth()
		if isHealthy {
			break
		}
		time.Sleep(time.Second)
	}

	if !isHealthy {
		err := s.DkrPool.Client.Logs(docker.LogsOptions{
			Container:    s.priceFeederResource.Container.ID,
			OutputStream: os.Stdout,
			ErrorStream:  os.Stderr,
			Stdout:       true,
			Stderr:       true,
			Tail:         "false",
		})
		if err != nil {
			s.T().Log("Error retrieving price feeder logs", err)
		}

		s.T().Fatal("price-feeder not healthy")
	}

	s.T().Logf("started price-feeder container: %s", s.priceFeederResource.Container.ID)
}

func (s *E2ETestSuite) connectIBCChains() {
	s.T().Logf("connecting %s and %s chains via IBC", s.Chain.ID, GaiaChainID)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	exec, err := s.DkrPool.Client.CreateExec(docker.CreateExecOptions{
		Context:      ctx,
		AttachStdout: true,
		AttachStderr: true,
		Container:    s.HermesResource.Container.ID,
		User:         "root",
		Cmd: []string{
			"hermes",
			"create",
			"channel",
			s.Chain.ID,
			GaiaChainID,
			"--port-a=transfer",
			"--port-b=transfer",
		},
	})
	s.Require().NoError(err)

	var (
		outBuf bytes.Buffer
		errBuf bytes.Buffer
	)

	err = s.DkrPool.Client.StartExec(exec.ID, docker.StartExecOptions{
		Context:      ctx,
		Detach:       false,
		OutputStream: &outBuf,
		ErrorStream:  &errBuf,
	})
	s.Require().NoErrorf(
		err,
		"failed connect chains; stdout: %s, stderr: %s", outBuf.String(), errBuf.String(),
	)

	s.Require().Containsf(
		errBuf.String(),
		"successfully opened init channel",
		"failed to connect chains via IBC: %s", errBuf.String(),
	)

	s.T().Logf("connected %s and %s chains via IBC", s.Chain.ID, GaiaChainID)
}

func noRestart(config *docker.HostConfig) {
	// in this case we don't want the nodes to restart on failure
	config.RestartPolicy = docker.RestartPolicy{
		Name: "no",
	}
}
