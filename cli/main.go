package main

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cgascoig/intersight-go/cli/commands"
	httptransport "github.com/go-openapi/runtime/client"
)

var (
	configFile string
)

func resultHandler(result interface{}, err error) {
	if err != nil {
		fmt.Printf("ERROR: %v", err)
		return
	}

	result = stripEnvelope(result)
	printResultHuman(result)
	// printResultYAML(result)
}

func main() {
	cobra.OnInitialize(initConfig)

	transport := httptransport.New("intersight.com", "/api/v1", []string{"https"})

	rootCmd := commands.GetCommands(transport, resultHandler)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.intersight-cli.yaml)")

	rootCmd.PersistentFlags().String("keyID", "", "API Key ID")
	rootCmd.PersistentFlags().String("keyFile", "", "API Private Key Filename")
	viper.BindPFlag("keyID", rootCmd.PersistentFlags().Lookup("keyID"))
	viper.BindPFlag("keyFile", rootCmd.PersistentFlags().Lookup("keyFile"))

	transport.DefaultAuthentication = authFunc()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".intersight-cli")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "WARNING: Can't read config:", err)
		// os.Exit(1)
	}
	// fmt.Printf("keyID: %s | %s\n", keyID, viper.GetString("keyID"))
}
