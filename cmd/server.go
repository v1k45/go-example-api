package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run shitposting API server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
