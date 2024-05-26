package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/v1k45/shitpost/api"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run shitposting API server",
	Run: func(cmd *cobra.Command, args []string) {
		addr := ":8080"

		server := api.NewServer(addr)
		fmt.Println("Server listening on ", addr)
		server.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
