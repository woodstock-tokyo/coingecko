package coingecko

import (
	"context"
	"testing"
)

func TestCoinList(t *testing.T) {
	client := NewClient()

	_, err := client.CoinList(context.Background(), nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
