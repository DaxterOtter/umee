package suite

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmrand "github.com/tendermint/tendermint/libs/rand"

	umeeapp "github.com/umee-network/umee/v6/app"
)

const (
	keyringPassphrase = "testpassphrase"
	keyringAppName    = "testnet"
)

var EncodingConfig sdkparams.EncodingConfig

func init() {
	EncodingConfig = umeeapp.MakeEncodingConfig()

	EncodingConfig.InterfaceRegistry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&stakingtypes.MsgCreateValidator{},
	)
	EncodingConfig.InterfaceRegistry.RegisterImplementations(
		(*cryptotypes.PubKey)(nil),
		&secp256k1.PubKey{},
		&ed25519.PubKey{},
	)
}

type SimpleChain struct {
	DataDir    string
	ID         string
	Validators []*Validator
}

func NewSimpleChain() (*SimpleChain, error) {
	tmpDir, err := os.MkdirTemp("", "umee-e2e-testnet-")
	if err != nil {
		return nil, err
	}

	return &SimpleChain{
		ID:      "Chain-" + tmrand.NewRand().Str(6),
		DataDir: tmpDir,
	}, nil
}

func (c *SimpleChain) ConfigDir() string {
	return fmt.Sprintf("%s/%s", c.DataDir, c.ID)
}

func (c *SimpleChain) CreateAndInitValidators(cdc codec.Codec, count int) error {
	for i := 0; i < count; i++ {
		node := c.CreateValidator(i)

		// generate genesis files
		if err := node.init(cdc); err != nil {
			return err
		}

		c.Validators = append(c.Validators, node)

		// create keys
		if err := node.createKey(cdc, "val"); err != nil {
			return err
		}
		if err := node.createNodeKey(); err != nil {
			return err
		}
		if err := node.createConsensusKey(); err != nil {
			return err
		}
	}

	return nil
}

func (c *SimpleChain) CreateValidator(index int) *Validator {
	return &Validator{
		Chain:   c,
		Index:   index,
		Moniker: "umee",
	}
}
