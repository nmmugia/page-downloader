package source

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

// LoggerFatal is to show log error to user followed by os.Exit(1) calling with message "[FATAL ERROR][NOW()] CAUSE()" as its format
func LoggerFatal(message string) {
	log.Fatalf("[FATAL ERROR][%s] Cause: %s", time.Now().Format(timeDefaultFormat), message)
}

// LoggerError is to show log error to user with "[ERROR][NOW()] CAUSE()" format
func LoggerError(message string) {
	log.Printf("[ERROR][%s] Cause: %s \n", time.Now().Format(timeDefaultFormat), message)
}

// LoggerInfo is to show log info to user with "[INFO][NOW()] CAUSE()" format
func LoggerInfo(message string) {
	log.Printf("[INFO][%s] %s \n", time.Now().Format(timeDefaultFormat), message)
}

// DownloadFile function is to download multifile concurrently
func DownloadFile(responsech chan *http.Response, errch chan error, urls []string) {
	for _, url := range urls {
		go func(url string) {
			response, err := http.Get(url)
			if err != nil {
				errch <- err
				responsech <- nil
				return
			}
			responsech <- response
			errch <- nil
		}(url)
	}
}

// getAttrValue function is to get the value attribute from a Token
func getAttrValue(n *html.Node, tag string, attr string, values *[]string) {
	if n.Type == html.ElementNode && n.Data == tag {
		for _, a := range n.Attr {
			if a.Key == attr {
				*values = append(*values, a.Val)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getAttrValue(c, tag, attr, values)
	}
}
