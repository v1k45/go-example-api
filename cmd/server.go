package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v1k45/shitpost/api"
	"github.com/v1k45/shitpost/config"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run shitposting API server",
	Run: func(cmd *cobra.Command, args []string) {
		server := api.NewServer(config.ServerAddr(), config.DatabaseUrl())
		fmt.Println("Server listening on ", config.ServerAddr())
		server.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
