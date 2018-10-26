package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

var rootCmd = &cobra.Command{
	Use:   "shard",
	Short: "Shard is a mythril light client",

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
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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