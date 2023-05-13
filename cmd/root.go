package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "jrpc",
	Short: "jrpc is a command line tool for generate JSON-RPC client and server code",
	Long:  `jrpc is a command line tool for generate JSON-RPC client and server code.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
