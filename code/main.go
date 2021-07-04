package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/nmmugia/page-downloader/code/source"
	src "github.com/nmmugia/page-downloader/code/source"
)

func main() {
	// guardian if argument isn't provided
	if len(os.Args) < 2 {
		src.LoggerFatal("No arguments provided")
	}

	var (
		withMetadata bool
		urls         []string
	)

	// filter args to get only argument and filter flag
	for i := 1; i < len(os.Args); i++ {
		switch {
		case os.Args[i] == "--metadata":
			withMetadata = true
		case strings.Index(os.Args[i], "http") == 0:
			urls = append(urls, os.Args[i])
		}
	}
	var (
		responsech = make(chan *http.Response, len(urls))
		errch      = make(chan error, len(urls))
	)
	src.DownloadFile(responsech, errch, urls)
	// download pages

	// main actions which are downloading page and log metadata if --metadata flag is exist
	for _, url := range urls {
		if err := <-errch; err != nil {
			src.LoggerError(err.Error())
		}
		var (
			data     bytes.Buffer
			response = <-responsech
		)

		_, err := io.Copy(&data, response.Body)
		if err != nil {
			src.LoggerError(err.Error())
			return
		}
		var (
			fetch = source.Init(url, data.Bytes())
		)
		if withMetadata {
			fetch.GetMetadataPage()
		}
		fetch.SavePageOffline()
	}
}
