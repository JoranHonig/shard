package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shard/core"
	"shard/mythril"
	"regexp"
)

var apiKey string
var analysisService core.AnalysisService

func init() {

	analyzeCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "The api key to authenticate with")
	viper.BindPFlag("api-key", analyzeCmd.Flags().Lookup("api-key"))
	rootCmd.AddCommand(analyzeCmd)
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze [bytecode|filename]",
	Short: "Analyzes the contract",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		verifyApiKey()
		setupAnalysisService()
	},
	Run: func(cmd *cobra.Command, args []string) {

		mode := determineMode(args[0])

		switch mode {
		case Bin:
			analyzeBytecode(args[0])
		case Filename:
			byteCodes, err := core.Compile(args[0])

			if err != nil {
				log.Fatal(err)
			}

			for _, contractcode := range byteCodes {
				analyzeBytecode(contractcode)
			}
		default:
			fmt.Println("Can't handle that input")
		}
	},
}

func setupAnalysisService() {
	s, err := mythril.BuildMythrilService(mythril.ALPHA, apiKey)
	if err != nil {
		log.Fatal(err)
	}
	analysisService = &core.BaseAnalysisService{MythrilService: s}
}

func verifyApiKey(){
	apiKey = viper.GetString("api-key")
	fmt.Println(apiKey)
	if len(apiKey) == 0 {
		println("No valid api key provided, exiting...")
		log.Exit(0)
	}
}

type InputType int
const (
	Bin InputType = 1 << iota
	Filename
	BinRuntime
)

func determineMode(argument string) (InputType) {
	isBytecode, err := regexp.MatchString("^(0x)?([0-9a-fA-F]{2})+$", argument)

	var inputType InputType

	switch {
	case isBytecode:
		inputType = Bin
	default:
		inputType = Filename
	}

	if err != nil {
		log.Fatal(err)
	}

	return inputType
}

func analyzeBytecode(bytecode string) {
	log.Info(fmt.Sprintf("Starting analysis for: %s", bytecode))
	_, err := analysisService.AnalyzeBytecode(bytecode)
	if err != nil {
		log.Fatal(err)
	}
}
