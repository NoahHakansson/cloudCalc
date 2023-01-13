// Package main
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var nrOfRequests = 500

func main() {
	// take user input
	fmt.Println("Please enter how many requests to send: ")
	fmt.Scanln(&nrOfRequests)
	fmt.Println("Thank you!")

	// do POST request
	fmt.Print("\nStarting requets...\n\n")
	url := "http://calc-business-logic-lb-7c9f93818ccf5e08.elb.us-east-1.amazonaws.com:5000/api/calc"

	payload := strings.NewReader("{\n\t\"first\": 2,\n\t\"second\": 1000,\n\t\"operator\": \"^\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	for i := 0; i < nrOfRequests-1; i++ {
		fmt.Println("requet NR: ", i+1)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(res)
		fmt.Println(string(body))
	}
}
