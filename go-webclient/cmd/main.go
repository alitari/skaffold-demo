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
	version := os.Getenv("VERSION")
	fmt.Printf("go-webclient version: %s : \n", version)
	urls := os.Getenv("URLS")
	urla := strings.Split(urls, ",")
	for counter := 0; ; counter++ {
		fmt.Printf("Attempt %d : \n", counter)
		for _, url := range urla {
			resp, err := http.Get(url)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Printf("GET call to endpoint %s\nbody:'%s'\n\n", url, body)
		}
		time.Sleep(time.Second * 1)
	}
}
