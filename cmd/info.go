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
	"log"
	"os"

	"github.com/pokom/go-wallpaper/providers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Featured struct {
	Image    string `yaml:"image"`
	Provider string `yaml:"provider"`
	Source   string `yaml:"source"`
	Title    string `yaml:"source"`
}

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print info on featured artwork.",
	Long:  `Fetch info on the featured artwork and print it out.`,
	Run: func(cmd *cobra.Command, args []string) {
		var featured Featured
		err := viper.UnmarshalKey("featured", &featured)
		if err != nil {
			log.Fatal(err)
		}

		err = providers.PrintTempl(os.Stdout, &providers.ImageResponse{
			ImageURI: featured.Image,
			Source:   featured.Source,
			Title:    featured.Title,
		})
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
