package coingecko

import (
	"context"
)

// CoinList coin list response
type CoinList []struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// CoinListOption coin list option
type CoinListOption struct {
	IncludePlatform bool `url:"include_platform,omitempty"`
}

// MarketChart List all supported coin id, name and symbol
func (c Client) CoinList(ctx context.Context, opt *CoinListOption) (CoinList, error) {
	cl := CoinList{}

	endpoint, err := c.endpointWithOpts("/coins/list", opt)
	if err != nil {
		return cl, err
	}

	err = c.GetJSON(ctx, endpoint, &cl)
	return cl, err
}
