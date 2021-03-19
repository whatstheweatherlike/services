package scrape

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strconv"
	"strings"
)

var Client *http.Client

func init() {
	Client = http.DefaultClient
}

type Location struct {
	Latitude, Longitude float64
}

// FIXME: this approach won't work together with AJAX requests, the elements we are searching for are not yet there
// at the time we are looking for them
func Scrape(address string) (results []string, err error) {
	var locationServiceURI = "https://www.ridemypark.com/explore-2/?search_keywords=%s"
	scrapeURI := fmt.Sprintf(locationServiceURI, address)
	resp, err := Client.Get(scrapeURI)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		defer resp.Body.Close()
		htmlTokenizer := html.NewTokenizer(resp.Body)
		for {
			token := htmlTokenizer.Next()
			switch token {
			case html.ErrorToken:
				return nil, htmlTokenizer.Err()
			case html.StartTagToken:
				t := htmlTokenizer.Token()
				spot := false
				var latitude, longitude float64
				for _, attr := range t.Attr {
					switch attr.Key {
					case "class":
						if strings.Contains(attr.Val, "type-spot") {
							spot = true
						}
					case "data-latitude":
						latitude, err = strconv.ParseFloat(attr.Val, 64)
						if err != nil {
							return nil, err
						}
						break
					case "data-longitude":
						longitude, err = strconv.ParseFloat(attr.Val, 64)
						if err != nil {
							return nil, err
						}
						break
					default:
						continue
					}
				}
			default:
				continue
			}
		}
	} else {
		return nil, fmt.Errorf("Status code %d from URI %s", resp.StatusCode, scrapeURI)
	}
}
