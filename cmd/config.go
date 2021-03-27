package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/ytakahashi/coas/internal"
)

var configFile string

type configFileContent struct {
	OasFiles []string
}

var config configFileContent

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.AddConfigPath("$HOME/.coas")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err == nil {
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func readConfig() string {
	if len(config.OasFiles) == 0 {
		return ""
	}
	selector := &internal.FuzzySelector{
		Items: config.OasFiles,
	}
	result, err := selector.Run()
	if err != nil {
		fmt.Printf("Failed to read input. %v\n", err)
		os.Exit(1)
	}
	return result
}
