###Warning! ALPHA software!

The dnode-client-go library is designed to make it easier for developers to send transactions and make queries to the dfinance blockchain network.

This library is still under heavy development and is not intended for production use.

The library now contains the minimum methods needed to sign and send transactions to the network. 
The full list of available methods will be available later.

### Example
A simple example of use: sending asset prices to the oracle module. 
To work with this example, you need to have dnode installed. 
The installation process is described in the document at: https://github.com/dfinance/dnode#installation.


After installation, just start the dnode daemon:
```$ dnode start``` and REST-server ```$ dncli rest-server```


```Go
package main

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	sdkcli "github.com/cosmos/cosmos-sdk/client"
	
	dncl"github.com/dfinance/dnode-client-go"
)

func main() {
	accountName := "your-account-name"
	// make Keybase from dncli context.
	kb, err = sdkcli.NewKeyBaseFromDir(os.ExpandEnv("$HOME/.dncli"))
	if err != nil {
		panic(err)
	}
	
	// make TxBuilder to sign transactions
	txb := auth.TxBuilder{}.
		WithKeybase(kb).
		WithChainID("your-chain-id").
		WithFees("your-fees").
		WithGas(200000)

	// make dnode client
	apiCl := dncl.New(
		dncl.WithTxBuilder(txb),
		dncl.WithAccountName(accountName),
		dncl.WithPassphrase("sccount-passphrase"),
	)
	
	// get your account information from the network
	keyInfo, err := kb.Get(accountName)
	if err != nil {
		panic(err)
	}
	acc, err := apiCl.Auth().Account(keyInfo.GetAddress())
	if err != nil {
		panic(err)
	}

	result, err := apiCl.WithAccount(acc).Oracle().PostPrices([]MsgPostPrice{
		{
			From:       ki.GetAddress(),
			AssetCode:  "eth_dfi",
			Price:      sdk.NewInt(1000000),
			ReceivedAt: time.Now(),
		},
		{
			From:       ki.GetAddress(),
			AssetCode:  "eth_dfi",
			Price:      sdk.NewInt(1200000),
			ReceivedAt: time.Now(),
		},
	})
}
```

###TODO

- [ ] A lot of tests 
- [ ] Methods available for calling but not yet implemented in this library
