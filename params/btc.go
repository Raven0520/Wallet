package params

// SegwitPath path parameters
type SegwitPath struct {
	Account  uint32 `json:"account" form:"account" validate:"-" comment:"Account Number"`
	External bool   `json:"external" form:"external" validate:"-" comment:"If External"`
	Address  uint32 `json:"address" form:"address" validate:"-" comment:"Address Number"`
}

// MnemonicPost Generate segwit from mnemonic parameters
type MnemonicPost struct {
	ChainID  uint64 `json:"chain_id" form:"chain_id" validate:"required" comment:"Blockchain ID"`
	Mnemonic string `json:"mnemonic" form:"mnemonic" validate:"required" comment:"Mnemonic"`
	Pass     string `json:"pass" form:"pass" validate:"-" comment:"Password"`
	SegwitPath
}

// SeedPost Generate segwit from seed parameters
type SeedPost struct {
	ChainID uint64 `json:"chain_id" form:"chain_id" validate:"required" comment:"Blockchain ID"`
	Seed    string `json:"seed" form:"seed" validate:"required,min=128,max=512" comment:"Seed"`
	SegwitPath
}

// MultiSigPost Generate Multiple signatures parameters
type MultiSigPost struct {
	ChainID    uint64   `json:"chain_id" form:"chain_id" validate:"required" comment:"Blockchain ID"`
	Required   uint64   `json:"required" form:"required" validate:"required" comment:"Number of signatures are required"`
	PublicKeys []string `json:"public_keys" form:"public_keys" validate:"required" comment:"Public Key Pool"`
}
