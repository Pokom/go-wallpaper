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

	"github.com/pokom/go-muzei/providers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		provider := viper.GetString("provider")
		client := providers.NewProvider(provider)
		featured, err := client.GetLatestImage()
		if err != nil {
			log.Fatal(err)
		}
		imageFileName := providers.BuildFileName(featured.ImageURI)
		fileName := filepath.Join(pictureDir, imageFileName)
		file := providers.CreateFile(fileName)
		err = client.DownloadImage(file, featured.ImageURI)
		if err != nil {
			log.Fatal(err)
		}

		itermBackgroundOsascript := fmt.Sprintf(`
tell application "iTerm2"
  tell current session of current window
	set background image to "%s"
  end tell
end tell`, fileName)

		command := exec.Command("osascript", "-e", itermBackgroundOsascript, fileName)
		_, err = command.Output()
		if err != nil {
			log.Fatal(err)
		}

		err = providers.PrintTempl(os.Stdout, featured)
		if err != nil {
			log.Fatal(err)
		}

		if err := providers.SaveFeatured(cmd.Name(), provider, featured); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(itermCmd)
}
