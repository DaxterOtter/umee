package keeper

import (
	"context"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/umee-network/umee/v6/x/oracle/types"
)

var _ types.QueryServer = querier{}

// Querier implements a QueryServer for the x/oracle module.
type querier struct {
	Keeper
}

// NewQuerier returns an implementation of the oracle QueryServer interface
// for the provided Keeper.
func NewQuerier(keeper Keeper) types.QueryServer {
	return &querier{Keeper: keeper}
}

// Params queries params of x/oracle module.
func (q querier) Params(
	goCtx context.Context,
	req *types.QueryParams,
) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// ExchangeRates queries exchange rates of all denoms, or, if specified, returns
// a single denom.
func (q querier) ExchangeRates(
	goCtx context.Context,
	req *types.QueryExchangeRates,
) (*types.QueryExchangeRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: need to decide if we want to return DecCoins here or list of ExchangeRates with denoms (we
	// need the latter for genesis anyway)
	var exchangeRates sdk.DecCoins

	if len(req.Denom) > 0 {
		exchangeRate, err := q.GetExchangeRate(ctx, req.Denom)
		if err != nil {
			return nil, err
		}

		exchangeRates = exchangeRates.Add(sdk.NewDecCoinFromDec(req.Denom, exchangeRate.Rate))
	} else {
		q.IterateExchangeRates(ctx, func(denom string, exgRate sdk.Dec, _ time.Time) (stop bool) {
			exchangeRates = exchangeRates.Add(sdk.NewDecCoinFromDec(denom, exgRate))
			return false
		})
	}

	return &types.QueryExchangeRatesResponse{ExchangeRates: exchangeRates}, nil
}

// ActiveExchangeRates queries all denoms for which exchange rates exist.
func (q querier) ActiveExchangeRates(
	goCtx context.Context,
	req *types.QueryActiveExchangeRates,
) (*types.QueryActiveExchangeRatesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	denoms := []string{}
	q.IterateExchangeRates(ctx, func(denom string, _ sdk.Dec, _ time.Time) (stop bool) {
		denoms = append(denoms, denom)
		return false
	})

	return &types.QueryActiveExchangeRatesResponse{ActiveRates: denoms}, nil
}

// FeederDelegation queries the account address to which the validator operator
// delegated oracle vote rights.
func (q querier) FeederDelegation(
	goCtx context.Context,
	req *types.QueryFeederDelegation,
) (*types.QueryFeederDelegationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	feederAddr, err := q.GetFeederDelegation(ctx, valAddr)
	if err != nil {
		return nil, err
	}

	return &types.QueryFeederDelegationResponse{
		FeederAddr: feederAddr.String(),
	}, nil
}

// MissCounter queries oracle miss counter of a validator.
func (q querier) MissCounter(
	goCtx context.Context,
	req *types.QueryMissCounter,
) (*types.QueryMissCounterResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryMissCounterResponse{
		MissCounter: q.GetMissCounter(ctx, valAddr),
	}, nil
}

// SlashWindow queries the current slash window progress of the oracle.
func (q querier) SlashWindow(
	goCtx context.Context,
	req *types.QuerySlashWindow,
) (*types.QuerySlashWindowResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params := q.GetParams(ctx)

	return &types.QuerySlashWindowResponse{
		WindowProgress: (uint64(ctx.BlockHeight()) % params.SlashWindow) /
			params.VotePeriod,
	}, nil
}

// AggregatePrevote queries an aggregate prevote of a validator.
func (q querier) AggregatePrevote(
	goCtx context.Context,
	req *types.QueryAggregatePrevote,
) (*types.QueryAggregatePrevoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	prevote, err := q.GetAggregateExchangeRatePrevote(ctx, valAddr)
	if err != nil {
		return nil, err
	}

	return &types.QueryAggregatePrevoteResponse{
		AggregatePrevote: prevote,
	}, nil
}

// AggregatePrevotes queries aggregate prevotes of all validators
func (q querier) AggregatePrevotes(
	goCtx context.Context,
	req *types.QueryAggregatePrevotes,
) (*types.QueryAggregatePrevotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var prevotes []types.AggregateExchangeRatePrevote
	q.IterateAggregateExchangeRatePrevotes(ctx, func(_ sdk.ValAddress, prevote types.AggregateExchangeRatePrevote) bool {
		prevotes = append(prevotes, prevote)
		return false
	})

	return &types.QueryAggregatePrevotesResponse{
		AggregatePrevotes: prevotes,
	}, nil
}

// AggregateVote queries an aggregate vote of a validator
func (q querier) AggregateVote(
	goCtx context.Context,
	req *types.QueryAggregateVote,
) (*types.QueryAggregateVoteResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	valAddr, err := sdk.ValAddressFromBech32(req.ValidatorAddr)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	vote, err := q.GetAggregateExchangeRateVote(ctx, valAddr)
	if err != nil {
		return nil, err
	}

	return &types.QueryAggregateVoteResponse{
		AggregateVote: vote,
	}, nil
}

// AggregateVotes queries aggregate votes of all validators
func (q querier) AggregateVotes(
	goCtx context.Context,
	req *types.QueryAggregateVotes,
) (*types.QueryAggregateVotesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var votes []types.AggregateExchangeRateVote
	q.IterateAggregateExchangeRateVotes(ctx, func(_ sdk.ValAddress, vote types.AggregateExchangeRateVote) bool {
		votes = append(votes, vote)
		return false
	})

	return &types.QueryAggregateVotesResponse{
		AggregateVotes: votes,
	}, nil
}

// Medians queries medians of all denoms, or, if specified, returns
// a single median.
func (q querier) Medians(
	goCtx context.Context,
	req *types.QueryMedians,
) (*types.QueryMediansResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	var medians types.Prices

	if len(req.Denom) > 0 {
		if req.NumStamps == 0 {
			return nil, status.Error(codes.InvalidArgument, "parameter NumStamps must be greater than 0")
		}

		if req.NumStamps > uint32(q.MaximumMedianStamps(ctx)) {
			req.NumStamps = uint32(q.MaximumMedianStamps(ctx))
		}

		medians = q.HistoricMedians(ctx, req.Denom, uint64(req.NumStamps))
	} else {
		medians = q.AllMedianPrices(ctx)
	}

	return &types.QueryMediansResponse{Medians: medians.Sort()}, nil
}

// MedianDeviations queries median deviations of all denoms, or, if specified, returns
// a single median deviation.
func (q querier) MedianDeviations(
	goCtx context.Context,
	req *types.QueryMedianDeviations,
) (*types.QueryMedianDeviationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	medianDeviations := types.Prices{}

	if len(req.Denom) > 0 {
		price, err := q.HistoricMedianDeviation(ctx, req.Denom)
		if err != nil {
			return nil, err
		}
		medianDeviations = append(medianDeviations, price)
	} else {
		medianDeviations = q.AllMedianDeviationPrices(ctx)
	}

	return &types.QueryMedianDeviationsResponse{MedianDeviations: medianDeviations.Sort()}, nil
}

// AvgPrice queries historic avg price for requested denom.
func (q querier) AvgPrice(
	goCtx context.Context,
	req *types.QueryAvgPrice,
) (*types.QueryAvgPriceResponse, error) {
	if req.Denom == "" {
		return nil, status.Error(codes.InvalidArgument, "denom must be defined")
	}
	if err := sdk.ValidateDenom(req.Denom); err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprint("malformed denom:", err))
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	p, err := q.HistoricAvgPrice(ctx, strings.ToUpper(req.Denom))
	if err != nil {
		return nil, err
	}
	return &types.QueryAvgPriceResponse{Price: p}, nil
}

// ExgRatesWithTimestamp queries exchange rates of all denoms with timestamp, or, if specified, returns
// a single denom.
func (q querier) ExgRatesWithTimestamp(
	goCtx context.Context,
	req *types.QueryExgRatesWithTimestamp,
) (*types.QueryExgRatesWithTimestampResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: need to decide if we want to return DecCoins here or list of ExchangeRates with denoms (we
	// need the latter for genesis anyway)
	var exgRates []types.DenomExchangeRate

	if len(req.Denom) > 0 {
		exchangeRate, err := q.GetExchangeRate(ctx, req.Denom)
		if err != nil {
			return nil, err
		}
		exgRates = append(exgRates, types.NewDenomExchangeRate(req.Denom, exchangeRate.Rate, exchangeRate.Timestamp))
	} else {
		q.IterateExchangeRates(ctx, func(denom string, exgRate sdk.Dec, t time.Time) (stop bool) {
			exgRates = append(exgRates, types.NewDenomExchangeRate(denom, exgRate, t))
			return false
		})
	}

	return &types.QueryExgRatesWithTimestampResponse{ExgRates: exgRates}, nil
}
