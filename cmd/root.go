package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "go-cobra-tot01",
	Short: "its my first cli with go ... lucifer's inthe house babe",
	Long: `f
  f
  f
  f should be long so what???
  f
  f
  f
  `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)

// 	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-cobra-tot01.yaml)")

// 	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }

// func initConfig() {
// 	if cfgFile != "" {
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".go-cobra-tot01")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
