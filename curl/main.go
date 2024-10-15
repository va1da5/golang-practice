package main

// go run main.go -u http://parrot.live
// go run main.go -u http://ipconfig.me
// go run main.go -u https://wttr.in

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var url string
var output string

func main() {
	flag.StringVar(&url, "u", "", "URL to get")
	flag.StringVar(&output, "o", "", "write output to file")
	flag.Parse()

	if len(url) < 4 {
		fmt.Println("Usage: curl -u <url>")
		os.Exit(0)
	}

	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	request.Header.Set("User-Agent", "curl/7.54.1")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if len(output) != 0 {
		body, err := io.ReadAll(response.Body)

		if err != nil {
			panic(err)
		}

		os.WriteFile(output, body, 0644)
		return
	}

	_, err = io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

}
