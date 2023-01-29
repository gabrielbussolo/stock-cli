package client

import (
	"context"
	"fmt"
	ab "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"os"
)

type Ticker struct {
	MarketCap float64
	Name      string
	Symbol    string
}

type polygon struct {
	client *ab.Client
}

func (p *polygon) GetTicker(ticker string) (*Ticker, error) {
	params := models.GetTickerDetailsParams{
		Ticker: ticker,
	}
	details, err := p.client.GetTickerDetails(context.Background(), &params)
	if err != nil {
		return nil, err
	}
	return &Ticker{
		Name:      details.Results.Name,
		Symbol:    details.Results.Ticker,
		MarketCap: details.Results.MarketCap,
	}, nil
}
func NewPolygon() (*polygon, error) {
	apiKey := os.Getenv("POLYGON_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("POLYGON_API_KEY is not set")
	}
	client := ab.New(apiKey)
	return &polygon{
		client: client,
	}, nil
}
