package suite

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	appparams "github.com/umee-network/umee/v6/app/params"
)

const (
	PhotonDenom    = "photon"
	InitBalanceStr = "510000000000" + appparams.BondDenom + ",100000000000" + PhotonDenom
)

var (
	minGasPrice     = appparams.ProtocolMinGasPrice.String()
	stakeAmount, _  = sdk.NewIntFromString("100000000000")
	stakeAmountCoin = sdk.NewCoin(appparams.BondDenom, stakeAmount)

	stakeAmount2, _  = sdk.NewIntFromString("500000000000")
	stakeAmountCoin2 = sdk.NewCoin(appparams.BondDenom, stakeAmount2)
)

var (
	ATOM          = "ATOM"
	ATOMBaseDenom = "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	ATOMExponent  = 6
)

func CreateMnemonic() (string, error) {
	entropySeed, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}

	mnemonic, err := bip39.NewMnemonic(entropySeed)
	if err != nil {
		return "", err
	}

	return mnemonic, nil
}
