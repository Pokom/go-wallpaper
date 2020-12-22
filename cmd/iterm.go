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
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pokom/iterm-muzei/providers"
	"github.com/spf13/cobra"
)

// itermCmd represents the iterm command
var itermCmd = &cobra.Command{
	Use:   "iterm",
	Short: "Update iterm's background image.",
	Long:  `Fetch the latest image from muzei and set iTerms background image to it.`,
	Run: func(cmd *cobra.Command, args []string) {
		pictureDir := os.Getenv("PICTURE_DIR")
		if len(pictureDir) == 0 {
			log.Fatal("PICTURE_DIR Needs to be set")
		}
		muzeiClient := providers.NewMuzeiClient()
		featured, err := muzeiClient.GetFeatured()
		if err != nil {
			log.Fatal(err)
		}
		imageFileName := providers.BuildFileName(featured.ImageURI)
		fileName := filepath.Join(pictureDir, imageFileName)
		file := providers.CreateFile(fileName)
		err = providers.DownloadImage(file, featured.ImageURI)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Wrote out image to: %s\n", fileName)
		command := exec.Command("bash", "set-iterm-img.sh", fileName)
		command.Dir = "bin"
		_, err = command.Output()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(itermCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// itermCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// itermCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
