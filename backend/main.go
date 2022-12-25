// Package main
package main

import (
	"log"

	"github.com/NoahHakansson/cloudCalc/backend/src/server"
)

func main() {
	r, err := server.Start()
	if err != nil {
		log.Println(err)
		return
	}
	err = r.Run(":5000")
	if err != nil {
		log.Println(err)
		return
	}
}
