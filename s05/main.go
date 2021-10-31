package main

/*
   @Author: Bhaskar S
   @Blog:   https://www.polarsparc.com
   @Date:   30 Oct 2021
*/

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	req, err := http.NewRequest("GET", "http://localhost:8080/", nil)
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
