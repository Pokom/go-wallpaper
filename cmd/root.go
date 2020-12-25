/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/pokom/go-muzei/providers"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "muzei",
	Short: "cli Muzei implementation. Dynamically set background image on iterm or wallpaper",
	Long: `Muzei is an Android application that is intended to help explore pieces of artwork currently
	hung up in museums. This app replicates some of Muzei's features by enabling the user to set the background image
	of iterm or the desktop.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.iterm-muzei.yaml)")
	rootCmd.PersistentFlags().String("provider", "muzei", "Provider to source artwork from. Defaults to muzei")
	rootCmd.PersistentFlags().String("subreddit", providers.SUBREDDIT, fmt.Sprintf("Default subreddit. Default is %s", providers.SUBREDDIT))
	rootCmd.PersistentFlags().Bool("random", true, "Choose random from top ten of subreddit. If false, choses first. Defaults to true")

	viper.BindPFlag("provider", rootCmd.PersistentFlags().Lookup("provider"))
	viper.BindPFlag("subreddit", rootCmd.PersistentFlags().Lookup("subreddit"))
	viper.BindPFlag("random", rootCmd.PersistentFlags().Lookup("random"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".iterm-muzei" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".iterm-muzei")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
