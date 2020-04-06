package client

import (
	"fmt"
)

// OracleClient provides a set of methods for the Oracle module
type OracleClient struct {
	dc DnodeClient
}

// PostPrices sends asset prices to the oracle module.
func (c OracleClient) PostPrices(msgs []MsgPostPrice) (result TxResponse, err error) {
	msg := make([]Msg, len(msgs))
	for i, p := range msgs {
		msg[i] = p
	}

	return c.dc.Tx().Broadcast(msg)
}

// Assets returns the array of available assets of the oracle module.
func (c OracleClient) Assets() (Assets, error) {
	resp, err := c.dc.httpcl.R().Get(fmt.Sprintf("%s/oracle/assets", c.dc.nodeURL))
	if err != nil {
		return Assets{}, err
	}

	var res = struct {
		Height string `json:"height"`
		Result Assets `json:"result"`
	}{}
	err = c.dc.httpcl.JSONUnmarshal(resp.Body(), &res)
	if err != nil {
		return Assets{}, err
	}

	return res.Result, nil
}
