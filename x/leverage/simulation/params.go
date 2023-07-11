package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/umee-network/umee/v5/x/leverage/types"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(*rand.Rand) []simtypes.LegacyParamChange {
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyCompleteLiquidationThreshold),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenCompleteLiquidationThreshold(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyMinimumCloseFactor),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenMinimumCloseFactor(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyOracleRewardFactor),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenOracleRewardFactor(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeySmallLiquidationSize),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenSmallLiquidationSize(r))
			},
		),
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.KeyDirectLiquidationFee),
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenDirectLiquidationFee(r))
			},
		),
	}
}
