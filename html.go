package html

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titulo obtêm o título de uma página HTML.
func Titulo(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			match := r.FindStringSubmatch(string(html))

			if len(match) < 1 {
				c <- fmt.Sprintf("Cannot get the HTML title for %s", url)
			} else {
				c <- match[1]
			}
		}(url)
	}
	return c
}
