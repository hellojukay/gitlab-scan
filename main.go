package main

import (
	"log"

	"github.com/hellojukay/gitlab/client"
)

func main() {
	client := client.New("", "")
	if !client.Ping() {
		log.Fatalln("can not connect to gitlab")
	}
}
