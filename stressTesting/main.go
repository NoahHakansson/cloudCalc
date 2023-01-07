package main

import (
	"github.com/gin-gonic/gin"
  "fmt"
)

func main() {
  shouldQuit := false
  address := ""
  port := ""
  times := 0
  fmt.Printf("IP-address of API-server: \n")
  fmt.Scanln(&address)
  fmt.Printf("Port of API-server: \n")
  fmt.Scanln(&port)
  for shouldQuit != true {
    fmt.Printf("Times to send exponent 2^1000 to API-server: \n")
    fmt.Scanln(&times)
    
  }
}
