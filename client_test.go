// +build integration

package client

import (
	"bufio"
	"os"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	crkeys "github.com/cosmos/cosmos-sdk/crypto/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/dfinance/dnode/cmd/config"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
)

var (
	DefaultCLIHome = os.ExpandEnv("$HOME/.dncli")
)

func init() {
	cfg := sdk.GetConfig()
	config.InitBechPrefixes(cfg)
	cfg.Seal()
}

func makeKeybaseAndClient(t *testing.T) (crkeys.Keybase, DnodeClient) {
	inBuf := bufio.NewReader(os.Stdin)
	kb, err := keys.NewKeyring("dfinance", "os", viper.GetString(flags.FlagHome), inBuf)
	require.NoError(t, err)

	txBuiler := auth.TxBuilder{}.
		WithKeybase(kb).
		WithChainID("dn-testnet").
		WithFees("1dfi").
		WithGas(500000)

	client := New(WithTxBuilder(txBuiler), WithAccountName("oracle1"))
	require.NoError(t, err)

	return kb, client
}

func TestDnodeClient_PostPrices(t *testing.T) {
	kb, cl := makeKeybaseAndClient(t)

	ki, err := kb.Get("oracle1")
	require.NoError(t, err)

	acc, err := cl.Auth().Account(ki.GetAddress())
	require.NoError(t, err)
	result, err := cl.WithAccount(acc).Oracle().PostPrices([]MsgPostPrice{
		{
			From:       ki.GetAddress(),
			AssetCode:  "eth_usdt",
			Price:      sdk.NewInt(1000000),
			ReceivedAt: time.Now(),
		},
		{
			From:       ki.GetAddress(),
			AssetCode:  "dfi_eth",
			Price:      sdk.NewInt(1200000),
			ReceivedAt: time.Now(),
		},
	})
	require.NoError(t, err)
	require.True(t, result.Code == 0)

	assets, err := cl.Oracle().Assets()
	require.NoError(t, err)
	require.True(t, len(assets) > 0)
}

func TestDnodeClient_IssueCurrency(t *testing.T) {
	kb, cl := makeKeybaseAndClient(t)

	ki, err := kb.Get("validator1")
	require.NoError(t, err)

	acc, err := cl.Auth().Account(ki.GetAddress())
	issueMsg := MsgIssueCurrency{
		ID:    "Issuing USDT",
		Coin:  sdk.NewCoin("USDT", sdk.NewInt(1000)),
		Payee: acc.GetAddress(),
	}
	resp, err := cl.WithAccount(acc).WithAccountName("validator1").WithBroadcastMode(TxBlockMode).Currencies().Issue(issueMsg, "msgID#1")
	require.NoError(t, err)
	require.True(t, resp.Code == 0)
}
