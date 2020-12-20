package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func main() {
	muzeiClient := NewMuzeiClient()
	resp, err := muzeiClient.GetFeatured()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response: %v", resp)
}
