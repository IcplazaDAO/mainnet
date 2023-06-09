package simulation

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/gauss/gauss/v6/x/token/types"
)

const (
	KeyCommunityTax      = "CommunityTax"
	KeyIssueTokenFee     = "IssueTokenFee"
	KeyMintTokenFeeRatio = "MintTokenFeeRatio"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, KeyCommunityTax,
			func(r *rand.Rand) string {
				return RandomDec(r).String()
			},
		),
		simulation.NewSimParamChange(types.ModuleName, KeyIssueTokenFee,
			func(r *rand.Rand) string {
				return RandomInt(r).String()
			},
		),
		simulation.NewSimParamChange(types.ModuleName, KeyMintTokenFeeRatio,
			func(r *rand.Rand) string {
				return RandomDec(r).String()
			},
		),
	}
}
