package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   30 Oct 2021
*/

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler - start ...")
	defer log.Println("indexHandler - done !!!")

	ctx := r.Context()

	go func() {
		dummyDbHandler(ctx)
	}()

	select {
	case <-ctx.Done():
		log.Printf("indexHandler - %v", ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusExpectationFailed)
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "<h3>Hello from Go Server !!!</h3>")
	}
}

func dummyDbHandler(ctx context.Context) {
	log.Println("dummyDbHandler - start ...")
	defer log.Println("dummyDbHandler - done !!!")

	select {
	case <-ctx.Done():
		log.Printf("dummyDbHandler - %v", ctx.Err())
	case <-time.After(5 * time.Second):
		log.Println("dummyDbHandler - completed DB operation ...")
	}
}

func main() {
	log.Println("Ready to start server on *:8080...")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
