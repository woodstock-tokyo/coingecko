package coingecko

import (
	"context"
	"fmt"
)

// Coins coins response
type Coins struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description struct {
		English  string `json:"en"`
		Japanese string `json:"ja"`
	} `json:"description,omitempty"`
	Links struct {
		HomePage []string `json:"homepage"`
	} `json:"links"`
	Image struct {
		Large string `json:"large"`
	} `json:"image"`
	MarketCapRank uint64 `json:"market_cap_rank"`
	MarketData    struct {
		TotalSupply       float64 `json:"total_supply"`
		MaxSupply         float64 `json:"max_supply"`
		CirculatingSupply float64 `json:"circulating_supply"`
	} `json:"market_data"`
}

// CoinsOption option for fetching coins
type CoinsOption struct {
	IncludeLocalizedLang bool `url:"localization,omitempty"`
	IncludeTickersData   bool `url:"tickers,omitempty"`
	IncludeMarketData    bool `url:"market_data,omitempty"`
	IncludeCommunityData bool `url:"community_data,omitempty"`
	IncludeDeveloperData bool `url:"developer_data,omitempty"`
	IncludeSparklineData bool `url:"sparkline,omitempty"`
}

// Coins get current data (name, price, market, ... including exchange tickers) for a coin
func (c Client) Coins(ctx context.Context, id string, opt *CoinsOption) (Coins, error) {
	coins := Coins{}
	endpoint, err := c.endpointWithOpts(fmt.Sprintf("/coins/%s", id), opt)
	if err != nil {
		return coins, err
	}

	err = c.GetJSON(ctx, endpoint, &coins)
	return coins, err
}
