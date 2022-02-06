package coingecko

import (
	"context"
	"fmt"
)

// MarketChart market chart response
type MarketChart struct {
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

// MarketChartOption market chart option
type MarketChartOption struct {
	TargetCurrency Currency `url:"vs_currency"`
	Days           int      `url:"days"`
	Interval       Interval `url:"interval,omitempty"`
}

// MarketChart Get historical market data include price, market cap, and 24h volume (granularity auto)
func (c Client) MarketChart(ctx context.Context, id string, opt *MarketChartOption) (MarketChart, error) {
	m := MarketChart{}
	if opt == nil {
		err := fmt.Errorf("must assign valid option")
		return m, err
	}

	endpoint, err := c.endpointWithOpts(fmt.Sprintf("/coins/%s/market_chart", id), opt)
	if err != nil {
		return m, err
	}

	err = c.GetJSON(ctx, endpoint, &m)
	return m, err
}
