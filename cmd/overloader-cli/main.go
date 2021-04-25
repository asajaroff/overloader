package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	//"time"
)

// handleRequests recives the URL, method and paylod (if request =! GET)
// and returns the HTTP response status code
func handleRequest(target string, method string) {
	switch method {
	case "GET":
		log.Println("Sending ", method, "requests to target (", target, ")")
		response, err := http.Get(target)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(response.StatusCode)
		log.Print(response.Body)

		bytes, errRead := ioutil.ReadAll(response.Body)
		if errRead != nil {
			log.Fatal(errRead)
		}
		log.Println("Response payload: ", len(string(bytes)), " bytes.")
		return

	case "HEAD":
		// ToDo
		log.Println("Implementation missing")

	case "POST":
		// ToDo
		log.Println("Implementation missing")

	case "PUT":
		// ToDo
		log.Println("Implementation missing")

	case "DELETE":
		// ToDo
		log.Println("Implementation missing")

	case "CONNECT":
		// ToDo
		log.Println("Implementation missing")

	case "OPTIONS":
		// ToDo
		log.Println("Implementation missing")

	case "TRACE":
		// ToDo
		log.Println("Implementation missing")

	case "PATCH":
		// ToDo
		log.Println("Implementation missing")
	}
}

const count = 500
const url = "https://www.google.com"

func main() {
	wg := sync.WaitGroup{}

	for id := 0; id <= count; id++ {
		wg.Add(1)
		go func(id int) {
			// response, err := handleRequest(url, "GET")
			// if err != nil {
			// 	return
			// }
			handleRequest(url, "GET")
			fmt.Println("Request ", id)

			wg.Done()

		}(id)
	}
}
