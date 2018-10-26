package cmd

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/compiler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"shard/mythril"
	"strings"
	"errors"
	"github.com/spf13/viper"
)

var contractBytecode string
var filename string
var apiKey string

func init() {

	analyzeCmd.Flags().StringVarP(&contractBytecode, "runtime-bytecode", "c", "", "The runtime bytecode to analyze")
	analyzeCmd.Flags().StringVarP(&filename, "filename", "f", "", "The contract to analyze")
	analyzeCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "The api key to authenticate with")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(analyzeCmd)
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyzes the contract",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(apiKey) == 0 {
			apiKey = viper.GetString("api-key")
		}
		if len(apiKey) == 0 {
			println("No valid api key provided, exiting...")
			log.Exit(0)
		}

		if len(filename) == 0 && len(contractBytecode) == 0 {
			println("You must provide either runtime bytecode or a filename")
			log.Exit(0)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		if contractBytecode != "" {
			log.Info(fmt.Sprintf("Starting analysis for: %s", contractBytecode))
			mythrilService, err := mythril.BuildMythrilService(mythril.ALPHA, apiKey)

			if err != nil {
				log.Fatal(err)
			}

			mythrilService.Submit(contractBytecode)
			return
		}

		byteCodes, err := compile(filename)

		if err != nil {
			log.Fatal(err)
		}

		for _, contractcode := range byteCodes {
			fmt.Println(contractcode)
			// TODO: analyze bytecode
		}
	},
}

// Compiles the contract at _filename
// _filename can also be of the form filename:ContractName
// In which case only the contract with ContractName will be considered
func compile(_filename string) ([]string, error) {
	parts := strings.Split(_filename, ":")

	if len(parts) > 1 {
		contracts, err := compiler.CompileSolidity("", parts[0])
		if err != nil {
			return nil, err
		}

		contract, ok := contracts[_filename]
		if !ok {
			return nil, errors.New("Wrong contract name provided")

		}
		return []string{contract.Code}, nil
	}

	contracts, err := compiler.CompileSolidity("", filename)
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	for _, contract := range contracts {
		result = append(result, contract.Code)
	}

	return result, nil
}
