package client

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dnode/x/currencies"
	"github.com/dfinance/dnode/x/multisig"
	"github.com/dfinance/dnode/x/oracle"
	"github.com/dfinance/dnode/x/poa"
)

// DefaultCodec returns the initialized amino codec to marshal/unmarshal
func DefaultCodec() *codec.Codec {
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	sdk.RegisterCodec(cdc)
	oracle.RegisterCodec(cdc)
	currencies.RegisterCodec(cdc)
	multisig.RegisterCodec(cdc)
	poa.RegisterCodec(cdc)

	return cdc
}
