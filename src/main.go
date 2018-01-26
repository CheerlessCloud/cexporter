package main

import (
	"fmt"
	"log"
	"net/http"
)

var config = GetConfig()

func main() {
	FetchMetrics(config.dockerURL)
	return
	fmt.Println(config)

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	})

	log.Fatal(http.ListenAndServe(config.httpHost+":"+config.httpPort, nil))
}
