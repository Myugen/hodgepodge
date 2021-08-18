package cmd

import (
	"log"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "utils",
		Short: "Script utils for stuffs",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath("/etc/utils")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("config.yml")
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Error execution of CLI. Cause: %s", err)
	}
}
