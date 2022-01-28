package types

const (
	ModuleName   = "auction"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
	MemStoreKey  = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
