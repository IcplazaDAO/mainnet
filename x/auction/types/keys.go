package types

const (
	// ModuleName defines the module name
	ModuleName = "auction"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_auction"

	// this line is used by starport scaffolding # ibc/keys/name
)

var (
	PrefixEndTime            = "End-time-"
	PrefixOrder              = "Order-"
	PrefixOrderTokenIdToPool = "TokenId-To-Pool-"
)

// this line is used by starport scaffolding # ibc/keys/port

func KeyPrefix(p string) []byte {
	return []byte(p)
}
