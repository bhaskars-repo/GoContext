package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   31 Oct 2021
*/

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("PS-TXN-ID")

	log.Printf("[%s] indexHandler - start ...", id)
	defer log.Printf("[%s] indexHandler - done !!!", id)

	ctx := context.WithValue(r.Context(), "PS-TXN-ID", id)

	go func() {
		dummyDbHandler(ctx)
	}()

	select {
	case <-ctx.Done():
		log.Printf("[%s] indexHandler - %v", id, ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusExpectationFailed)
	case <-time.After(time.Second):
		fmt.Fprintln(w, "<h3>Hello from Go Server !!!</h3>")
	}
}

func dummyDbHandler(ctx context.Context) {
	id := ctx.Value("PS-TXN-ID")

	log.Printf("[%s] dummyDbHandler - start ...", id)
	defer log.Printf("[%s] dummyDbHandler - done !!!", id)

	select {
	case <-ctx.Done():
		log.Printf("[%s] dummyDbHandler - %v", id, ctx.Err())
	case <-time.After(time.Second):
		log.Printf("[%s] dummyDbHandler - completed DB operation ...", id)
	}
}

func main() {
	log.Println("Ready to start server on *:8080...")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
