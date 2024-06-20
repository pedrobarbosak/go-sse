package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pedrobarbosak/go-sse"
)

func main() {
	port := 8080

	sse := sse.New(sse.NewConfig())
	event := "event-name" // can be get by param, url, etc

	http.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		sse.Upgrade(w, r, event)
	})

	// Produce events
	go func() {
		count := 0

		for {
			time.Sleep(time.Second * 3)

			count++

			if err := sse.PublishJSON(event, count); err != nil {
				log.Println("failed to publish:", err)
			}
		}
	}()

	// Start server
	fmt.Println("Starting server on :", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("error starting server: %s\n", err)
	}
}
