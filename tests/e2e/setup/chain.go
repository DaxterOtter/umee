package setup

import (
	"github.com/umee-network/umee/v6/tests/suite"

	"github.com/cosmos/cosmos-sdk/codec"
)

const (
	keyringPassphrase = "testpassphrase"
	keyringAppName    = "testnet"
)

type e2eChain struct {
	*suite.SimpleChain
	GaiaValidators []*gaiaValidator
}

func newChain() (*e2eChain, error) {
	simpleChain, err := suite.NewSimpleChain()
	if err != nil {
		return nil, err
	}
	return &e2eChain{
		SimpleChain: simpleChain,
	}, nil
}

func (c *e2eChain) createAndInitGaiaValidator(cdc codec.Codec) error {
	// create gaia validator
	gaiaVal := c.createGaiaValidator(0)

	// create keys
	mnemonic, info, err := createMemoryKey(cdc)
	if err != nil {
		return err
	}

	gaiaVal.keyInfo = *info
	gaiaVal.mnemonic = mnemonic

	c.GaiaValidators = append(c.GaiaValidators, gaiaVal)

	return nil
}

func (c *e2eChain) createGaiaValidator(index int) *gaiaValidator {
	return &gaiaValidator{
		index: index,
	}
}
