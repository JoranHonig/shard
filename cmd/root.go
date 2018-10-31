package cmd

import (
	"fmt"
	"github.com/JoranHonig/shard/core"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

	"github.com/mitchellh/go-homedir"
)

var analysisService core.AnalysisService

var (
	RootCmd = &cobra.Command{
		Use:   "shard",
		Short: "Shard is a mythril light client",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())

			viper.BindPFlag("verbose", cmd.Flags().Lookup("verbose"))

			if viper.GetBool("verbose") {
				logrus.SetLevel(logrus.DebugLevel)
			} else {
				logrus.SetLevel(logrus.ErrorLevel)
			}
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Shard",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Shard v0.0.1")
		},
	}
	cfgFile string
)

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringP( "api-key", "k", "", "The api key to authenticate with. Overrides config value.")
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/.shard.yaml)")
	RootCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging.")
	viper.BindPFlag("api-key", analyzeCmd.Flags().Lookup("api-key"))

	RootCmd.AddCommand(versionCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	viper.SetConfigType("yaml")
	println(cfgFile)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home + "/.config")
		viper.AddConfigPath(home)
		viper.SetConfigName(".shard")
	}

	if err := viper.ReadInConfig(); err != nil {
		println("aaah")
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Info("Failed to load configuration")
	}
}
