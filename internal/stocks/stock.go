package stocks

import "gabrielbussolo/stock-cli/internal/platform/client"

type Service interface {
	GetTicker(ticker string) (*client.Ticker, error)
}

type stock struct {
	client client.Client
}

func New(c client.Client) *stock {
	return &stock{client: c}
}

func (s *stock) GetTicker(ticker string) (*client.Ticker, error) {
	return s.client.GetTicker(ticker)
}
