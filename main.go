package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const muzeiApi string = "https://muzeiapi.appspot.com/featured?cachebust=1"

type MuzeiResponse struct {
	Attribution string    `json:"attribution"`
	Byline      string    `json:"byline"`
	DetailsURI  string    `json:"detailsUri"`
	ImageURI    string    `json:"imageUri"`
	NextTime    time.Time `json:"nextTime"`
	ThumbURI    string    `json:"thumbUri"`
	Title       string    `json:"title"`
}

type MuzeiClient struct {
	Client      *http.Client
	FeaturedURL string
}

// NewMuzeiClient is a builder for a MuzeiClient
func NewMuzeiClient() *MuzeiClient {
	return &MuzeiClient{
		Client:      &http.Client{},
		FeaturedURL: muzeiApi,
	}
}

// GetFeatured will fetch the latest featured image from Muzei
func (mc *MuzeiClient) GetFeatured() (*MuzeiResponse, error) {
	resp, err := mc.Client.Get(mc.FeaturedURL)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Getting json failed: %s", resp.Status)
	}

	var result MuzeiResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DownloadImage will fetch a given image
func DownloadImage(file *os.File, imageURI string) error {
	fmt.Printf("Fetching image: %s\n", imageURI)
	resp, err := http.Get(imageURI)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Error retrieving image=%s, %s", imageURI, err)
	}

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		return err
	}
	fmt.Printf("Downloaded image=%s is %d", imageURI, size)
	return nil
}

func buildFileName(imageURI string) string {
	url, err := url.Parse(imageURI)

	if err != nil {
		log.Fatal(err)
	}
	path := url.Path
	segment := strings.Split(path, "/")
	return segment[len(segment)-1]
}

func createFile(imageURI string) *os.File {
	imageFileName := buildFileName(imageURI)
	file, err := os.Create(imageFileName)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func main() {
	muzeiClient := NewMuzeiClient()
	featured, err := muzeiClient.GetFeatured()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Inject prefix for where to store image
	file := createFile(featured.ImageURI)
	err = DownloadImage(file, featured.ImageURI)
	if err != nil {
		log.Fatal(err)
	}
	// TODO: Execute osascript to set the downloaded image as the background
}
