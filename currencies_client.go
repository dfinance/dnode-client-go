package client

// CurrenciesClient provides a set of methods for the Currencies module
type CurrenciesClient struct {
	dc DnodeClient
}

// Issue currency
func (c CurrenciesClient) Issue(msg MsgIssueCurrency, msgID string) (TxResponse, error) {
	return c.dc.Tx().Broadcast([]Msg{Msg(MsgSubmitCall{Msg: msg, UniqueID: msgID, Creator: c.dc.fromAddress})})
}

// Withdraw currency
func (c CurrenciesClient) Withdraw(msg MsgWithdrawCurrency) (TxResponse, error) {
	return c.dc.Tx().Broadcast([]Msg{msg})
}
