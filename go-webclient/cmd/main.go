package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := os.Getenv("URLS")
	version := os.Getenv("VERSION")
	for counter := 0; ; counter++ {
		urla := strings.Split(urls, ",")
		for _, url := range urla {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("(%v) version: %s - GET call from %s = \n%s", counter, version, url, body)
		}
		time.Sleep(time.Second * 1)
	}
}
