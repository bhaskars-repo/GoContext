package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   30 Oct 2021
*/

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://httpbin.org/delay/5", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("HTTP Status code: %d", res.StatusCode)
}
