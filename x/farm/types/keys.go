package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// nolint
const (
	// module name
	ModuleName = "farm"

	// StoreKey is the default store key for farm
	StoreKey = ModuleName

	// RouterKey is the message route for farm
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the farm store.
	QuerierRoute = StoreKey

	// Query endpoints supported by the farm querier
	QueryRecord = "farm"

	// RewardCollector is the root string for the reward distribution account address
	RewardCollector = "reward_collector"
)

var (
	FarmPoolKey       = []byte{0x01} // key for farm pool
	FarmPoolRuleKey   = []byte{0x02} // key for farm pool reward rule
	FarmerKey         = []byte{0x03} // key for farmer
	ActiveFarmPoolKey = []byte{0x04} // key for active farm pool
	FarmersRewardsKey = []byte{0x05} // key for farmers rewards
	FarmerRewardsKey  = []byte{0x06} // key for farmer rewards

	// Separator for string key
	Delimiter = []byte{0x00}
)

func KeyFarmPool(poolName string) []byte {
	return append(FarmPoolKey, []byte(poolName)...)
}

func KeyRewardRule(poolName, reward string) []byte {
	key := append(FarmPoolRuleKey, []byte(poolName)...)
	return append(append(key, Delimiter...), []byte(reward)...)
}

func PrefixRewardRule(poolName string) []byte {
	key := append(FarmPoolRuleKey, []byte(poolName)...)
	return append(key, Delimiter...)
}

func PrefixFarmersRewards(poolName string) []byte {
	key := append(FarmersRewardsKey, []byte(poolName)...)
	return append(key, Delimiter...)
}
func PrefixFarmerRewards(poolName string, sender string) []byte {
	key := append(append(FarmerRewardsKey, []byte(poolName)...), []byte(sender)...)
	return append(key, Delimiter...)
}

func KeyFarmerInfo(address, poolName string) []byte {
	key := append(FarmerKey, []byte(address)...)
	key = append(key, Delimiter...)
	key = append(key, []byte(poolName)...)
	return key
}

func PrefixFarmInfo(address string) []byte {
	return append(FarmerKey, []byte(address)...)
}

func KeyActiveFarmPool(height int64, poolName string) []byte {
	return append(append(ActiveFarmPoolKey, sdk.Uint64ToBigEndian(uint64(height))...), []byte(poolName)...)
}

func PrefixActiveFarmPool(height int64) []byte {
	return append(ActiveFarmPoolKey, sdk.Uint64ToBigEndian(uint64(height))...)
}
