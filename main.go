package main

import (
	"log"
	"os"

	"github.com/hellojukay/gitlab/client"
)

func main() {
	var api = os.Getenv("API")
	client := client.New(api, os.Getenv("TOKEN"))
	if !client.Ping() {
		log.Fatalln("can not connect to gitlab")
	} else {
		log.Printf("auth %s success", api)
	}
}
