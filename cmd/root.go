package cmd

import (
	"fmt"
	"gabrielbussolo/stock-cli/cmd/get"
	"gabrielbussolo/stock-cli/internal/platform/client"
	"gabrielbussolo/stock-cli/internal/stocks"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "stock-cli",
	Short: "A CLI to get stock information",
	Long:  `A CLI to get stock information from the Polygon.io API`,
}

func Execute() {
	polygon, err := client.NewPolygon()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	stock := stocks.New(polygon)
	cmdGet := get.NewCmdGet(stock)
	rootCmd.AddCommand(cmdGet)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
