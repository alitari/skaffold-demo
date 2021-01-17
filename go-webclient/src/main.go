package main

import (
	"fmt"
	"time"
	"os"
	"net/http"
	"io/ioutil"
)

func main() {
	url := os.Getenv("URL")
	version := os.Getenv("VERSION")
	for counter := 0; ; counter++ {
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Printf("(%v) version: %s - GET call from %s = \n%s", counter, version,url,body)
		time.Sleep(time.Second * 1)
	}
}
