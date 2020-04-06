package client

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	resty "github.com/go-resty/resty/v2"
)

type Option func(*DnodeClient)

// WithHTTPClient sets the http.Client for the DnodeClient instance being created instead of the default one
func WithHTTPClient(cl *http.Client) Option {
	return func(c *DnodeClient) {
		c.httpcl = resty.NewWithClient(cl)
	}
}

// WithNodeURL sets the node url for the DnodeClient instance being created instead of the default one
func WithNodeURL(url string) Option {
	return func(c *DnodeClient) {
		c.nodeURL = url
	}
}

// WithCodec sets the codec for the DnodeClient instance being created instead of the default one
func WithCodec(cdc *codec.Codec) Option {
	return func(c *DnodeClient) {
		c.cdc = cdc
	}
}

// WithTxBuilder sets the transactions builder for the DnodeClient instance being created
func WithTxBuilder(b auth.TxBuilder) Option {
	return func(c *DnodeClient) {
		c.txb = b
	}
}

// WithAccountName sets the transaction sender's account name for the DnodeClient instance being created
func WithAccountName(name string) Option {
	return func(c *DnodeClient) {
		c.accName = name
	}
}

// WithPassphrase sets the passphrase for the DnodeClient instance being created
func WithPassphrase(pp string) Option {
	return func(c *DnodeClient) {
		c.passphrase = pp
	}
}

// WithBroadcastMode sets the transactions broadcast mode for the DnodeClient instance being created instead of the default one
func WithBroadcastMode(mode TxBroadcastMode) Option {
	return func(c *DnodeClient) {
		c.broadcastMode = mode.String()
	}
}

// WithFromAddress sets the transaction sender's address for the DnodeClient instance being created
func WithFromAddress(addr AccAddress) Option {
	return func(c *DnodeClient) {
		c.fromAddress = addr
	}
}
