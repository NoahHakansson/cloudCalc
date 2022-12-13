// Package main
package main

import (
	"log"

	"github.com/NoahHakansson/cloudCalc/backend/src/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Println(err)
		return
	}
}
