package get

import (
	"gabrielbussolo/stock-cli/internal/stocks"
	"github.com/spf13/cobra"
	"golang.org/x/text/currency"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func NewCmdGet(stock stocks.Service) *cobra.Command {
	var flagVerbose bool
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a stock",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if flagVerbose {
				cmd.Printf("Verbose output enabled")
			}
			ticker, err := stock.GetTicker(args[0])
			if err != nil {
				cmd.Printf("Error: %s", err)
				return
			}

			pt := language.AmericanEnglish
			p := message.NewPrinter(pt)
			eur := currency.USD
			amt := eur.Amount(ticker.MarketCap)
			sprint := p.Sprint(currency.Symbol(amt))

			cmd.Printf("Market Cap for %s: %s", ticker.Symbol, sprint)
		},
	}
	cmd.Flags().BoolVarP(&flagVerbose, "verbose", "v", false, "Verbose output")
	return cmd
}
