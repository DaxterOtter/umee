package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/umee-network/umee/v5/x/leverage/types"
)

// minimum borrow factor is the minimum collateral weight and minimum liquidation threshold
// allowed when a borrowed token is limiting the efficiency of a pair of assets.
// TODO: parameterize this in the leverage module
var minimumBorrowFactor = sdk.MustNewDecFromStr("0.5")

// GetAccountPosition creates and sorts an accountPosition for an address, using information
// from the keeper's special asset pairs and token collateral weights as well as oracle prices.
// Will treat collateral with missing prices as zero-valued, but will error on missing borrow prices.
// On computing liquidation threshold, will treat borrows with missing prices as zero and error on
// missing collateral prices instead, as well as using spot prices instead of both spot and historic.
// Also stores all token settings and any special asset pairs that could apply to the account's collateral.
func (k Keeper) GetAccountPosition(ctx sdk.Context, addr sdk.AccAddress, isForLiquidation bool,
) (types.AccountPosition, error) {
	tokenSettings := k.GetAllRegisteredTokens(ctx)
	specialPairs := []types.SpecialAssetPair{
		// Simulating the presence of special pairs on mainnet state
		{
			// stATOM -> ATOM
			Collateral: "ibc/C8B3026C2844D204F3A49F91058DC947F01F2FC36AFF17850FFC8701504BDDEE",
			Borrow: "ibc/C4CFF46FD6DE35CA4CF4CE031E643C8FDC9BA4B99AE598E9B0ED98FE3A2319F9",
			CollateralWeight: sdk.MustNewDecFromStr("0.85"),
			LiquidationThreshold: sdk.MustNewDecFromStr("0.9"),
		},
		{
			// stOSMO -> OSMO
			Collateral: "ibc/6B49A789937D4E50BF01F0F50DDEDF5C1103EDF01306B7021BDF23BDE65D99BA",
			Borrow: "ibc/ED07A3391A112B175915CD8FAF43A2DA8E4790EDE12566649D0C2F97716B8518",
			CollateralWeight: sdk.MustNewDecFromStr("0.75"),
			LiquidationThreshold: sdk.MustNewDecFromStr("0.8"),
		},
	}
	collateral := k.GetBorrowerCollateral(ctx, addr)
	collateralValue := sdk.NewDecCoins()
	borrowed := k.GetBorrowerBorrows(ctx, addr)
	borrowedValue := sdk.NewDecCoins()

	var (
		v   sdk.Dec
		err error
	)

	// get the borrower's collateral value by token
	for _, c := range collateral {
		if isForLiquidation {
			// for liquidation threshold, error on collateral without prices
			// and use spot price
			v, err = k.CalculateCollateralValue(ctx, sdk.NewCoins(c), types.PriceModeSpot)
		} else {
			// for borrow limit, max borrow, and max withdraw, ignore collateral without prices
			// and use the lower of historic or spot prices
			v, err = k.VisibleCollateralValue(ctx, sdk.NewCoins(c), types.PriceModeLow)
		}
		if err != nil {
			return types.AccountPosition{}, err
		}
		denom := types.ToTokenDenom(c.Denom)
		collateralValue = collateralValue.Add(sdk.NewDecCoinFromDec(denom, v))
		// get special asset pairs which could apply to this collateral token
		// specialPairs = append(specialPairs, k.GetSpecialAssetPairs(ctx, c.Denom)...)
	}

	// get the borrower's borrowed value by token
	for _, b := range borrowed {
		if isForLiquidation {
			// for liquidation threshold, ignore borrow without prices
			// and use spot price
			v, err = k.VisibleTokenValue(ctx, sdk.NewCoins(b), types.PriceModeSpot)
		} else {
			// for borrow limit, max borrow, and max withdraw, error on borrow without prices
			// and use the higher of historic or spot prices
			v, err = k.TokenValue(ctx, b, types.PriceModeHigh)
		}
		if err != nil {
			return types.AccountPosition{}, err
		}
		if v.IsPositive() {
			borrowedValue = borrowedValue.Add(sdk.NewDecCoinFromDec(b.Denom, v))
		}
	}

	return types.NewAccountPosition(
		tokenSettings, specialPairs, collateralValue, borrowedValue, isForLiquidation, minimumBorrowFactor,
	)
}
