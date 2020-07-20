package client

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	curr "github.com/dfinance/dnode/x/currencies"
	msmsgs "github.com/dfinance/dnode/x/multisig"
	"github.com/dfinance/dnode/x/oracle"
)

type (
	// TxResponse defines a structure containing relevant tx data and metadata. The
	// tags are stringified and the log is JSON decoded.
	TxResponse = types.TxResponse
	// AccAddress a wrapper around bytes meant to represent an account address.
	// When marshaled to a string or JSON, it uses Bech32.
	AccAddress = types.AccAddress
	// BaseAccount - a base account structure.
	// This can be extended by embedding within in your AppAccount.
	// However one doesn't have to use BaseAccount as long as your struct
	// implements Account.
	BaseAccount = auth.BaseAccount
	// Msg transactions messages must fulfill the Msg
	Msg = types.Msg
	// MsgPostPrice struct representing a posted price message.
	// Used by oracles to input prices to the oracle
	MsgPostPrice = oracle.MsgPostPrice
	// MsgIssueCurrency struct for issue new currencies.
	// IssueID could be txHash of transaction in another blockchain.
	MsgIssueCurrency = curr.MsgIssueCurrency
	// MsgWithdrawCurrency message for destroy currency
	MsgWithdrawCurrency = curr.MsgWithdrawCurrency
	// Assets array type for oracle module
	Assets = oracle.Assets
	// MsgSubmitCall message for submit call
	MsgSubmitCall = msmsgs.MsgSubmitCall
	// MsgConfirmCall message for confirm call
	MsgConfirmCall = msmsgs.MsgConfirmCall
	// MsgRevokeConfirm message to revoke confirmation from call
	MsgRevokeConfirm = msmsgs.MsgRevokeConfirm
	// TxBuilder implements a transaction context created in SDK modules.
	TxBuilder = auth.TxBuilder
	// Codec amino codec to marshal/unmarshal
	Codec = codec.Codec
	// Keybase exposes operations on a generic keystore
	Keybase = keys.Keybase
	// StdFee includes the amount of coins paid in fees and the maximum
	// gas to be used by the transaction. The ratio yields an effective "gasprice",
	// which must be above some miminum to be accepted into the mempool.
	StdFee = auth.StdFee
	// StdTx is a standard way to wrap a Msg with Fee and Signatures.
	// NOTE: the first signature is the fee payer (Signatures must not be nil).
	StdTx = auth.StdTx
	// BroadcastReq defines a tx broadcasting request.
	BroadcastReq = rest.BroadcastReq
)

var (
	// SortJSON takes any JSON and returns it sorted by keys. Also, all white-spaces
	// are removed.
	// This method can be used to canonicalize JSON to be returned by GetSignBytes,
	// e.g. for the ledger integration.
	// If the passed JSON isn't valid it will return an error.
	SortJSON = types.SortJSON
)

const (
	BroadcastBlock = flags.BroadcastBlock
	BroadcastSync  = flags.BroadcastSync
	BroadcastAsync = flags.BroadcastAsync
)

// TxBroadcastMode enumeration type for BroadcastMode numeric presentation
type TxBroadcastMode int

const (
	TxBlockMode TxBroadcastMode = iota
	TxSyncMode
	TxAsyncMode
)

func (bm TxBroadcastMode) String() string {
	return [...]string{BroadcastBlock, BroadcastSync, BroadcastAsync}[bm]
}
