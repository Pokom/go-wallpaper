package providers

import (
	"io"
	"log"
	"os"
)


type ImageResponse struct {
	Title, ImageURI, Source string
}

type Provider interface {
	// TODO: Define interface
	// TODO: Refactor GetFeatured of MuzeiClient
	GetLatestImage() (*ImageResponse, error)
	DownloadImage(*os.File, string) error
	PrintTempl(io.Writer, *ImageResponse) error
}

func NewProvider(provider string) Provider {
	switch provider {
	case "muzei":
		return NewMuzeiClient()
	case "reddit":
		return NewRedditClient()
	default:
		log.Fatal("not implemented")
	}
	return nil
}