package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/v1k45/shitpost/api"
	"github.com/v1k45/shitpost/config"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run shitposting API server",
	Run: func(cmd *cobra.Command, args []string) {
		server := api.NewServer(config.ServerAddr(), config.DatabaseUrl())
		slog.Info("server_starting", "addr", config.ServerAddr())
		server.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
