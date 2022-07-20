package query

import (
	"encoding/json"
	"fmt"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	lvtypes "github.com/umee-network/umee/v2/x/leverage/types"
	octypes "github.com/umee-network/umee/v2/x/oracle/types"
)

// AssignedQuery defines the query to be called.
type AssignedQuery uint16

const (
	// AssignedQueryBorrowed represents the call to query the Borrowed coins of an address.
	AssignedQueryBorrowed AssignedQuery = iota + 1
	// AssignedQueryExchangeRates represents the call to query the exchange rates
	// of all denoms.
	AssignedQueryExchangeRates
	// AssignedQueryRegisteredTokens represents the call of leverage get all registered tokens.
	AssignedQueryRegisteredTokens
	// AssignedQueryLeverageParams represents the call of the x/leverage module's parameters.
	AssignedQueryLeverageParams
	// AssignedQueryBorrowedValue represents the call to query the Borrowed amount of an
	// specific coin of an address.
	AssignedQueryBorrowedValue
	// AssignedQuerySupplied represents the call to query the Supplied amount of an address.
	AssignedQuerySupplied
	// AssignedQuerySuppliedValue represents the call to query the Supplied amount of an
	// address in USD.
	AssignedQuerySuppliedValue
	// AssignedQueryAvailableBorrow represents the call to query the Available
	// amount of an denom.
	AssignedQueryAvailableBorrow
	// AssignedQueryBorrowAPY represents the call to query the current borrow interest
	// rate on a token denom.
	AssignedQueryBorrowAPY
	// AssignedQuerySupplyAPY represents the call to query and derives the current supply
	// interest rate on a token denom.
	AssignedQuerySupplyAPY
	// AssignedQueryTotalSuppliedValue represents the call to query the market size of
	// an token denom in USD.
	AssignedQueryTotalSuppliedValue
	// AssignedQueryTotalSupplied represents the call to query the market size of
	// an token denom.
	AssignedQueryTotalSupplied
	// AssignedQueryReserveAmount represents the call to query the gets the amount
	// reserved of a specified token.
	AssignedQueryReserveAmount
	// AssignedQueryCollateral represents the call to query the collateral amount
	// of a user by token denomination. If the denomination is not supplied, all
	// of the user's collateral tokens are returned.
	AssignedQueryCollateral
	// AssignedQueryCollateralValue represents the call to query the total USD
	// value of a user's collateral, or the USD value held as a given base
	// asset's associated uToken denomination.
	AssignedQueryCollateralValue
	// AssignedQueryExchangeRate represents the call to query and calculate the
	// token:uToken exchange rate of a base token denom.
	AssignedQueryExchangeRate
	// AssignedQueryBorrowLimit represents the call to query and calculate the
	// borrow limit (in USD).
	AssignedQueryBorrowLimit
	// AssignedQueryLiquidationThreshold represents the call to query and calculate
	// the maximum borrowed value (in USD) that a borrower with given
	// collateral could reach before being eligible for liquidation.
	AssignedQueryLiquidationThreshold
	// AssignedQueryLiquidationTargets represents the call to query the list of
	// borrower addresses eligible for liquidation.
	AssignedQueryLiquidationTargets
	// AssignedQueryMarketSummary represents the call to query the market
	// summary data of an denom.
	AssignedQueryMarketSummary
	// AssignedQueryTotalCollateral represents the call to query the total collateral
	// system-wide of a given uToken denomination.
	AssignedQueryTotalCollateral
	// AssignedQueryTotalBorrowed represents the call to query the total borrowed
	// system-wide of a given token denomination.
	AssignedQueryTotalBorrowed
	// AssignedQueryActiveExchangeRates represents the call to query all active denoms.
	AssignedQueryActiveExchangeRates
	// AssignedQueryActiveFeederDelegation represents the call to query all the feeder
	// delegation of a validator.
	AssignedQueryActiveFeederDelegation
	// AssignedQueryMissCounter represents the call to query all the oracle
	// miss counter of a validator.
	AssignedQueryMissCounter
	// AssignedQueryAggregatePrevote represents the call to query an aggregate prevote of
	// a validator.
	AssignedQueryAggregatePrevote
	// AssignedQueryAggregatePrevotes represents the call to query an aggregate prevote of
	// all validators.
	AssignedQueryAggregatePrevotes
	// AssignedQueryAggregateVote represents the call to query an aggregate vote of
	// a validator.
	AssignedQueryAggregateVote
	// AssignedQueryAggregateVotes represents the call to query an aggregate vote of
	// all validators.
	AssignedQueryAggregateVotes
	// AssignedQueryOracleParams represents the call of the x/leverage module's
	// parameters.
	AssignedQueryOracleParams
)

// MarshalResponse marshals any response.
func MarshalResponse(resp interface{}) ([]byte, error) {
	bz, err := json.Marshal(resp)
	if err != nil {
		return nil, wasmvmtypes.UnsupportedRequest{Kind: fmt.Sprintf("error %+v umee query response error on marshal", err)}
	}
	return bz, err
}

// UmeeQuery wraps all the queries availables for cosmwasm smartcontracts.
type UmeeQuery struct {
	// Mandatory field to determine which query to call.
	AssignedQuery AssignedQuery `json:"assigned_query"`
	// Used to query the Borrowed coins of an address.
	Borrowed *lvtypes.QueryBorrowed `json:"borrowed,omitempty"`
	// Used to get the exchange rates of all denoms.
	ExchangeRates *octypes.QueryExchangeRates `json:"exchange_rates,omitempty"`
	// Used to query all the registered tokens.
	RegisteredTokens *lvtypes.QueryRegisteredTokens `json:"registered_tokens,omitempty"`
	// Used to query the x/leverage module's parameters.
	LeverageParams *lvtypes.QueryParams `json:"leverage_params,omitempty"`
	// Used to query an specific borrow address value in usd.
	BorrowedValue *lvtypes.QueryBorrowedValue `json:"borrowed_value,omitempty"`
	// Used to query an the amount Supplied of an address.
	Supplied *lvtypes.QuerySupplied `json:"supplied,omitempty"`
	// Used to query an the amount Supplied of an address in USD.
	SuppliedValue *lvtypes.QuerySuppliedValue `json:"supplied_value,omitempty"`
	// Used to query an the amount available to borrow.
	AvailableBorrow *lvtypes.QueryAvailableBorrow `json:"available_borrow,omitempty"`
	// Used to query an current borrow interest rate on a token denom.
	BorrowAPY *lvtypes.QueryBorrowAPY `json:"borrow_apy,omitempty"`
	// Used to derives the current supply interest rate on a token denom.
	SupplyAPY *lvtypes.QuerySupplyAPY `json:"supply_apy,omitempty"`
	// Used to get the market size in USD of a token denom.
	TotalSuppliedValue *lvtypes.QueryTotalSuppliedValue `json:"total_supplied_value,omitempty"`
	// Used to get the market size of a token denom.
	TotalSupplied *lvtypes.QueryTotalSupplied `json:"total_supplied,omitempty"`
	// Used to gets the amount reserved of a specified token.
	ReserveAmount *lvtypes.QueryReserveAmount `json:"reserve_amount,omitempty"`
	// Used to gets the collateral amount of a user by token denomination.
	// If the denomination is not supplied, all of the user's collateral tokens
	// are returned.
	Collateral *lvtypes.QueryCollateral `json:"collateral,omitempty"`
	// Used to gets the total USD value of a user's collateral, or
	// the USD value held as a given base asset's associated uToken denomination.
	CollateralValue *lvtypes.QueryCollateralValue `json:"collateral_value,omitempty"`
	// Used to calculate the token:uToken exchange rate of a base token denom.
	ExchangeRate *lvtypes.QueryExchangeRate `json:"exchange_rate,omitempty"`
	// Uses the price oracle to determine the borrow limit (in USD).
	BorrowLimit *lvtypes.QueryBorrowLimit `json:"borrow_limit,omitempty"`
	// determines the maximum borrowed value (in USD) that a borrower with given
	// collateral could reach before being eligible for liquidation.
	LiquidationThreshold *lvtypes.QueryLiquidationThreshold `json:"liquidation_threshold,omitempty"`
	// request to return a list of borrower addresses eligible for liquidation.
	LiquidationTargets *lvtypes.QueryLiquidationTargets `json:"liquidation_targets,omitempty"`
	// Used to get the summary data of an denom.
	MarketSummary *lvtypes.QueryMarketSummary `json:"market_summary,omitempty"`
	// Used to get the total collateral system-wide of a given uToken denomination.
	TotalCollateral *lvtypes.QueryTotalCollateral `json:"total_collateral,omitempty"`
	// Used to get the total borrowed system-wide of a given token denomination.
	TotalBorrowed *lvtypes.QueryTotalBorrowed `json:"total_borrowed,omitempty"`
	// Used to get all active denoms.
	ActiveExchangeRates *octypes.QueryActiveExchangeRates `json:"active_exchange_rates,omitempty"`
	// Used to get all feeder delegation of a validator.
	FeederDelegation *octypes.QueryFeederDelegation `json:"feeder_delegation,omitempty"`
	// Used to get all the oracle miss counter of a validator.
	MissCounter *octypes.QueryMissCounter `json:"miss_counter,omitempty"`
	// Used to get an aggregate prevote of a validator.
	AggregatePrevote *octypes.QueryAggregatePrevote `json:"aggregate_prevote,omitempty"`
	// Used to get an aggregate prevote of all validators.
	AggregatePrevotes *octypes.QueryAggregatePrevotes `json:"aggregate_prevotes,omitempty"`
	// Used to get an aggregate vote of a validator.
	AggregateVote *octypes.QueryAggregateVote `json:"aggregate_vote,omitempty"`
	// Used to get an aggregate vote of all validators.
	AggregateVotes *octypes.QueryAggregateVotes `json:"aggregate_votes,omitempty"`
	// Used to query the x/oracle module's parameters.
	OracleParams *octypes.QueryParams `json:"oracle_params,omitempty"`
}
