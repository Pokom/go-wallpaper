/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
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
)

// wallpaperCmd represents the wallpaper command
var wallpaperCmd = &cobra.Command{
	Use:   "wallpaper",
	Short: "Update Wallpapers background image",
	Long:  `Fetch the latest image from muzei and set desktop background image to use it.`,
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
		err = muzeiClient.DownloadImage(file, featured.ImageURI)
		if err != nil {
			log.Fatal(err)
		}

		osascript := fmt.Sprintf(`
tell application "System Events"
	tell every desktop
		set picture to "%s"
	end tell
end tell
		`, fileName)

		command := exec.Command("osascript", "-e", osascript)
		_, err = command.Output()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(wallpaperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wallpaperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wallpaperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
