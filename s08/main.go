package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   30 Oct 2021
*/

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	dur := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, dur)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Printf("HTTP content: %s", data)
}
