package coingecko

// Currency curreny
type Currency string

// to be added
const USD Currency = "usd"
const JPY Currency = "jpy"
const EUR Currency = "eur"

// Interval used for market chart (only daily supported)
type Interval string

const Daily Interval = "daily"
