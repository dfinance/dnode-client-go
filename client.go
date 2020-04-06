package client

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	resty "github.com/go-resty/resty/v2"
)

// DnodeClient implements set of methods for working with dnode blockchain.
// Never create an instance with the DnodeClient{} literal!
// To create a correctly initialized instance, always use the New() function.
type DnodeClient struct {
	httpcl        *resty.Client
	nodeURL       string
	cdc           *Codec
	txb           TxBuilder
	accName       string
	fromAddress   AccAddress
	passphrase    string
	broadcastMode string
}

// New is a constructor function for DnodeClient
func New(opts ...Option) DnodeClient {
	cl := DnodeClient{
		cdc:           DefaultCodec(),
		nodeURL:       "http://127.0.0.1:1317",
		httpcl:        resty.New(),
		broadcastMode: BroadcastBlock,
		txb:           auth.TxBuilder{},
	}

	for _, opt := range opts {
		opt(&cl)
	}

	return cl
}

// WithAccount returns DnodeClient with AccountNumber, Sequence and FromAddress fields set.
// This is a short record of the call of the following methods WithAccountNumber(...).WithSequence(...).WithFromAddress(...).
func (c DnodeClient) WithAccount(account BaseAccount) DnodeClient {
	c.txb = c.txb.WithAccountNumber(account.GetAccountNumber()).WithSequence(account.GetSequence())
	c.fromAddress = account.GetAddress()
	return c
}

// WithAccountNumber returns DnodeClient with the specified AccountNumber
func (c DnodeClient) WithAccountNumber(number uint64) DnodeClient {
	c.txb = c.txb.WithAccountNumber(number)
	return c
}

// WithChainID returns DnodeClient with the specified ChainID
func (c DnodeClient) WithChainID(chainID string) DnodeClient {
	c.txb = c.txb.WithChainID(chainID)
	return c
}

// WithFees returns DnodeClient with the specified Fees
func (c DnodeClient) WithFees(fees string) DnodeClient {
	c.txb = c.txb.WithFees(fees)
	return c
}

// WithGas returns DnodeClient with the specified gas
func (c DnodeClient) WithGas(gas uint64) DnodeClient {
	c.txb = c.txb.WithGas(gas)
	return c
}

// WithGasPrices returns DnodeClient with the specified gas prices
func (c DnodeClient) WithGasPrices(gasPrices string) DnodeClient {
	c.txb = c.txb.WithGasPrices(gasPrices)
	return c
}

// WithKeybase returns DnodeClient with the specified KeyBase
func (c DnodeClient) WithKeybase(kb Keybase) DnodeClient {
	c.txb = c.txb.WithKeybase(kb)
	return c
}

// WithMemo returns DnodeClient with the specified memo
func (c DnodeClient) WithMemo(memo string) DnodeClient {
	c.txb = c.txb.WithMemo(memo)
	return c
}

// WithSequence returns DnodeClient with the specified account sequence
func (c DnodeClient) WithSequence(seq uint64) DnodeClient {
	c.txb = c.txb.WithSequence(seq)
	return c
}

// WithAccountName returns DnodeClient with the specified account name
func (c DnodeClient) WithAccountName(name string) DnodeClient {
	c.accName = name
	return c
}

// WithPassphrase returns DnodeClient with the specified passphrase
func (c DnodeClient) WithPassphrase(phrase string) DnodeClient {
	c.passphrase = phrase
	return c
}

// WithBroadcastMode returns DnodeClient with the specified broadcast mode
func (c DnodeClient) WithBroadcastMode(mode TxBroadcastMode) DnodeClient {
	c.broadcastMode = mode.String()
	return c
}

// WithFromAddress returns DnodeClient with the specified transaction sender address
func (c DnodeClient) WithFromAddress(addr AccAddress) DnodeClient {
	c.fromAddress = addr
	return c
}

// Auth returns initialized AuthClient
func (c DnodeClient) Auth() AuthClient {
	return AuthClient{dc: c}
}

// Currencies returns initialized CurrenciesClient
func (c DnodeClient) Currencies() CurrenciesClient {
	return CurrenciesClient{dc: c}
}

// Oracle returns initialized OracleClient
func (c DnodeClient) Oracle() OracleClient {
	return OracleClient{dc: c}
}

// Tx returns initialized TxClient
func (c DnodeClient) Tx() TxClient {
	return TxClient{dc: c}
}

// BroadcastMode returns the current broadcast mode
func (c DnodeClient) BroadcastMode() string {
	return c.broadcastMode
}

// Passphrase returns the current passphrase
func (c DnodeClient) Passphrase() string {
	return c.passphrase
}

// FromAddress returns the current transaction sender address
func (c DnodeClient) FromAddress() AccAddress {
	return c.fromAddress
}

// AccountName returns the current transaction sender account name
func (c DnodeClient) AccountName() string {
	return c.accName
}

// NodeURL returns the current node URL
func (c DnodeClient) NodeURL() string {
	return c.nodeURL
}

// Codec returns the current codec
func (c DnodeClient) Codec() *Codec {
	return c.cdc
}

// TxBuilder returns the current transactions bulder
func (c DnodeClient) TxBuilder() TxBuilder {
	return c.txb
}
