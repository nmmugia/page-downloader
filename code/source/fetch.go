package source

import (
	"bytes"
	"fmt"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/html"
)

type Fetch struct {
	URL  *url.URL
	Page []byte
}

func Init(link string, page []byte) Fetch {
	u, err := url.Parse(link)
	if err != nil {
		LoggerFatal(err.Error())
	}
	if err = os.MkdirAll(u.Hostname(), os.ModePerm); err != nil {
		LoggerFatal(err.Error())
	}
	return Fetch{
		URL:  u,
		Page: page,
	}
}

// GetMetadataPage function is to show the metadata of fetched page such as total images, link, and last_fetched
func (fetch Fetch) GetMetadataPage() {
	fmt.Printf("site: %s\n", fetch.URL)
	var (
		links, img []string
	)
	doc, err := html.Parse(bytes.NewReader(fetch.Page))
	if err != nil {
		LoggerError(err.Error())
	}

	getAttrValue(doc, "a", "href", &links)
	getAttrValue(doc, "img", "src", &img)

	fmt.Printf("num_links: %d\n", len(links))
	fmt.Printf("images: %d\n", len(img))
	fmt.Printf("last_fetch: %s\n", time.Now().Format(timeUIFriendlyFormat))
}

// SavePageOffline function is to show the metadata of fetched page such as total images, link, and last_fetched
func (fetch Fetch) SavePageOffline() {
	fileWriter, err := os.Create(fmt.Sprintf("./%s/index.html", fetch.URL.Hostname()))
	if err != nil {
		LoggerError(err.Error())
		return
	}
	defer fileWriter.Close()
	_, err = fileWriter.Write(fetch.Page)
	if err != nil {
		LoggerError(err.Error())
		return
	}
}
