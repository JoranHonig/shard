package cmd

import (
	"github.com/spf13/cobra"
	"shard/mythril"
	log "github.com/sirupsen/logrus"
	"fmt"
)

var contractBytecode string
var apiKey string

func init() {

	analyzeCmd.Flags().StringVarP(&contractBytecode, "runtime-bytecode", "c", "00", "The runtime bytecode to execute")
	analyzeCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "The api key to authenticate with")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(analyzeCmd)
}

var analyzeCmd = &cobra.Command{
	Use: "analyze ",
	Short: "Analyzes the contract",
	Run: func(cmd *cobra.Command, args []string) {
		if len(apiKey) == 0 {
			println("No valid api key provided, exiting...")
			return
		}
		log.Info(fmt.Sprintf("Starting analysis for: %s", contractBytecode))
		mythrilService, err := mythril.BuildMythrilService(mythril.ALPHA, apiKey)

		if err != nil {
			log.Fatal(err)
		}

		mythrilService.Submit(contractBytecode)
	},
}

