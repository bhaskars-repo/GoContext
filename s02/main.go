package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   30 Oct 2021
*/

import (
	"bufio"
	"context"
	"log"
	"net/http"
	"os"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	req, err := http.NewRequestWithContext(ctx, "GET", "http://httpbin.org/delay/5", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	go func() {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadLine()
		log.Println("Ready to cancel request...")
		cancel()
	}()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("HTTP Status code: %d", res.StatusCode)
}
