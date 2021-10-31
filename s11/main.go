package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   31 Oct 2021
*/

import (
	"context"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func makeHttpRequest(ch chan bool) {
	ctx := context.Background()
	dur := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, dur)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	id := uuid.New()
	req.Header.Set("PS-TXN-ID", id.String())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[%s] %v", id, err)
		os.Exit(1)
	}

	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("[%s] %v", id, err)
		os.Exit(1)
	}

	log.Printf("[%s] HTTP content: %s", id, data)

	ch <- true
}

func main() {
	ch := make(chan bool)

	for i := 1; i <= 3; i++ {
		go makeHttpRequest(ch)
	}

	for i := 1; i <= 3; i++ {
		<-ch
	}
}
