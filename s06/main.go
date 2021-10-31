package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   30 Oct 2021
*/

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler - start ...")
	defer log.Println("indexHandler - done !!!")

	ctx := r.Context()

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		http.Error(w, ctx.Err().Error(), http.StatusExpectationFailed)
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "<h3>Hello from Go Server !!!</h3>")
	}
}

func main() {
	log.Println("Ready to start server on *:8080...")

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
