package client

import (
	"errors"
	"fmt"
	"net/http"
)

// TxClient provides a set of methods for working with transactions.
type TxClient struct {
	dc DnodeClient
}

// Broadcast signs and sends/broadcasts transactions to the network
func (c TxClient) Broadcast(msgs []Msg) (TxResponse, error) {
	var res TxResponse

	body, err := c.buildSignedBroadcastRequest(msgs)
	if err != nil {
		return res, err
	}
	resp, err := c.dc.httpcl.R().SetBody(body).Post(fmt.Sprintf("%s/txs", c.dc.nodeURL))
	if err != nil {
		return res, err
	}

	if resp.StatusCode() != http.StatusOK {
		err = errors.New(resp.String())
		return res, err
	}

	err = c.dc.cdc.UnmarshalJSON(resp.Body(), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c TxClient) buildSignedBroadcastRequest(msgs []Msg) ([]byte, error) {
	signedTx, err := c.dc.txb.SignStdTx(c.dc.accName, c.dc.passphrase, StdTx{
		Msgs: msgs,
		Fee:  StdFee{Amount: c.dc.txb.Fees(), Gas: c.dc.txb.Gas()},
		Memo: c.dc.txb.Memo(),
	}, false)
	if err != nil {
		return nil, err
	}

	bz, err := c.dc.cdc.MarshalJSON(BroadcastReq{Mode: c.dc.broadcastMode, Tx: signedTx})
	if err != nil {
		return nil, err
	}

	return SortJSON(bz)
}
