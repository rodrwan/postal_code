package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/yhat/scrape"

	"golang.org/x/net/html"
)

const (
	// BaseURL ...
	BaseURL = "https://codigopostal.correos.cl/?calle=%s&numero=%s&comuna=%s"
)

// Crawler visit BaseURL and make a request to crawl DOM looking for specific data.
// This method return a new html.Parse instance with the document root, to traverse
// specific node.
func Crawler(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: Failed to crawl \"" + url + "\"")
		return nil, err
	}

	root, err := html.Parse(resp.Body)
	if err != nil {
		if err != nil {
			fmt.Println("ERROR: Failed to parse \"" + url + "\"")
			return nil, err
		}
	}

	return root, nil
}

// PostalCode search "to_codigo" class that contain postal code.
func PostalCode(url string) string {
	root, err := Crawler(url)
	if err != nil {
		return ""
	}
	link, ok := scrape.Find(root, scrape.ByClass("tu_codigo"))
	if ok {
		return scrape.Text(link)
	}
	return ""
}

func main() {
	args := os.Args[1:]

	calle := url.QueryEscape(args[0])
	numero := url.QueryEscape(args[1])
	comuna := url.QueryEscape(args[2])
	mailURL := fmt.Sprintf(BaseURL, calle, numero, comuna)
	fmt.Println(PostalCode(mailURL))
}
