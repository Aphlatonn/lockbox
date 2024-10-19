package main

import (
	"lockbox/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
