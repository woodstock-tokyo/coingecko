package coingecko

import (
	"context"
	"testing"
)

func TestMarketChart(t *testing.T) {
	client := NewClient()

	opt := &MarketChartOption{
		TargetCurrency: USD,
		Days:           1,
		Interval:       Daily,
	}

	_, err := client.MarketChart(context.Background(), "bitcoin", opt)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
