package providers

import (
	"io"
	"log"
	"os"
	"text/template"

	"github.com/spf13/viper"
)

const InfoTmpl = `Title: {{.Title}}
Image: {{.ImageURI}}
Source: {{.Source}}
`

type ImageResponse struct {
	Title, ImageURI, Source string
}

type Provider interface {
	// TODO: Define interface
	// TODO: Refactor GetFeatured of MuzeiClient
	GetLatestImage() (*ImageResponse, error)
	DownloadImage(*os.File, string) error
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

func PrintTempl(dst io.Writer, res *ImageResponse) error {
	tmpl := template.Must(template.New("artwork").Parse(InfoTmpl))
	err := tmpl.Execute(dst, res)
	return err
}

func SaveFeatured(target string, provider string, featured *ImageResponse) error {
	viper.Set("featured.provider", provider)
	viper.Set("featured.source", featured.Source)
	viper.Set("featured.image", featured.ImageURI)
	viper.Set("featured.title", featured.Title)
	if err := viper.WriteConfig(); err != nil {
		return err
	}
	return nil
}
