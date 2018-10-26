package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "shard",
	Short: "shard is a mythril light client",

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Shard",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Shard v0.0.1")
	},
}

func Execute() {
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}