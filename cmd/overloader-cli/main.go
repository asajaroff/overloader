package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"net/http"

	//"time"
	//"fmt"
	toml "github.com/pelletier/go-toml"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func main() {
	config, err := toml.LoadFile("config.toml")

	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		target := "https://" + config.Get("target.domain").(string)
		log.Println("Target present in config.toml: ", target)
		resp, err := http.Get(target)
		if err != nil {
			log.Println("Error requesting ", target, ": ", err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: ", err)
		}
		fmt.Print(body)
		defer resp.Body.Close()
		// ...
	}
}
