package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"shard/core"
	"regexp"
	"shard/mythril/generic"
	"text/template"
	"os"
	"shard/mythril"
)

var apiKey string
var analysisService core.AnalysisService

func init() {
	analyzeCmd.Flags().StringVarP(&apiKey, "api-key", "k", "", "The api key to authenticate with. Overrides config value.")
	viper.BindPFlag("api-key", analyzeCmd.Flags().Lookup("api-key"))
	RootCmd.AddCommand(analyzeCmd)
}

var analyzeCmd = &cobra.Command{
	Use:   "analyze [bytecode|filename]",
	Short: "Analyzes the contract",
	Args:  cobra.MinimumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		// Check api key
		apiKey = viper.GetString("api-key")
		if len(apiKey) == 0 {
			println("No valid api key provided, exiting...")
			log.Exit(0)
		}

		// Setup mythril service
		s, err := mythril.BuildMythrilService(mythril.ALPHA, apiKey)
		if err != nil {
			log.Fatal(err)
		}
		analysisService = &core.BaseAnalysisService{MythrilService: s}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mode := determineMode(args[0])

		switch mode {
		case Bin:
			analyzeBytecode(args[0])
		case Filename:
			log.Fatal("Compilation not fully supported")

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
	issues, err := analysisService.AnalyzeBytecode(bytecode)
	if err != nil {
		log.Fatal(err)
	}
	if len(issues) == 0 {
		fmt.Println("No issues found")
	} else {
		printLeThings(issues)
	}
}

func printLeThings(issues []generic.Issue) {
	templ, err := template.New("IssueTemplate").Parse(
		"== {{.Title}} ==\n" +
			"Function: {{.Function}} \n" +
			"Type: {{.Type}} \n" +
			"Description: \n" +
			"{{.Description}}")
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range issues {
		templ.Execute(os.Stdout, issue)
	}
}
