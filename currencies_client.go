package client

// CurrenciesClient provides a set of methods for the Currencies module
type CurrenciesClient struct {
	dc DnodeClient
}

// Issue currency
func (c CurrenciesClient) Issue(msg MsgIssueCurrency, msgID string) (TxResponse, error) {
	return c.dc.Tx().Broadcast([]Msg{Msg(NewMsgSubmitCall(msg, msgID, c.dc.fromAddress))})
}

// Destroy currency
func (c CurrenciesClient) Destroy(msg MsgDestroyCurrency) (TxResponse, error) {
	return c.dc.Tx().Broadcast([]Msg{msg})
}
