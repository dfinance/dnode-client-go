package client

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

// AuthClient provides a set of methods for the Auth module
type AuthClient struct {
	dc DnodeClient
}

// Account returns account information including AccountNumber, Sequence and Address.
// The AccountNumber and Sequence fields are required to sign transactions.
func (c AuthClient) Account(address AccAddress) (auth.BaseAccount, error) {
	resp, err := c.dc.httpcl.R().Get(fmt.Sprintf("%s/auth/accounts/%s", c.dc.nodeURL, address))
	if err != nil {
		return auth.BaseAccount{}, err
	}

	var rd rest.ResponseWithHeight
	err = c.dc.cdc.UnmarshalJSON(resp.Body(), &rd)
	if err != nil {
		return auth.BaseAccount{}, err
	}
	var acc = struct {
		Type  string           `json:"type"`
		Value auth.BaseAccount `json:"value"`
	}{}
	err = c.dc.cdc.UnmarshalJSON(rd.Result, &acc)
	if err != nil {
		return auth.BaseAccount{}, err
	}

	return acc.Value, err
}
