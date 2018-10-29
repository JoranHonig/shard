package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

var RootCmd = &cobra.Command{
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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Shard",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Shard v0.0.1")
	},

}

func Execute() {
	setupViper()
	RootCmd.AddCommand(versionCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(setupViper)

	RootCmd.Flags().String("config", "", "config file (default is $HOME/.config/shard.yaml)")
	RootCmd.PersistentFlags().BoolP("verbose","v", false, "Enable verbose logging.")
}

func setupViper(){
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	viper.SetConfigName("shard") // name of config file (without extension)
	viper.AddConfigPath("$HOME/.config")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		logrus.Info(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}