package params

// Dictionary Parameter dictionary
type Dictionary struct {
	BTC
}

// BTC BTC Parameter dictionary
type BTC struct {
	MnemonicPost
	SeedPost
	MultiSigPost
}
