module github.com/dfinance/dnode-client-go

go 1.14

replace (
	github.com/cosmos/cosmos-sdk => github.com/dfinance/cosmos-sdk v0.38.4-0.20200924011713-eecea3b95cd5
)

require (
	github.com/cosmos/cosmos-sdk v0.39.0
	github.com/dfinance/dnode v0.6.1-0.20200924090157-d6bbddf8eb4e
	github.com/go-resty/resty/v2 v2.2.0
	github.com/spf13/viper v1.6.3
	github.com/stretchr/testify v1.6.1
)
