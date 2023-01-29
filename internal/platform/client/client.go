package client

type Client interface {
	GetTicker(ticker string) (*Ticker, error)
}
