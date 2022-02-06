package coingecko

import (
	"context"
	"testing"
)

func TestCoins(t *testing.T) {
	client := NewClient()

	opt := &CoinsOption{
		IncludeLocalizedLang: false,
		IncludeTickersData:   false,
		IncludeCommunityData: false,
		IncludeMarketData:    false,
		IncludeDeveloperData: false,
		IncludeSparklineData: false,
	}

	_, err := client.Coins(context.Background(), "bitcoin", opt)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
