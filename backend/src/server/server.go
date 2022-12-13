// Package server provides server
package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// variables

var errNotSupportedOperator = errors.New("Error: Operator is not supported")

// structs

type calcRequest struct {
	First    float64 `form:"first" xml:"first" json:"first"`
	Second   float64 `form:"second" xml:"second" json:"second"`
	Operator string  `form:"operator" xml:"operator" json:"operator"`
}

// Start function
//
// Starts server or returns an error.
func Start() error {
	r := gin.Default()
	// set trusted proxy to localhost
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return err
	}

	// routes
	r.POST("/api/calculate", calcEndpoint)

	if err = r.Run(":5000"); err != nil {
		return err
	}
	return nil
}

func calcEndpoint(c *gin.Context) {
	calcReq := calcRequest{}
	// bind body data or return error if it fails
	if err := c.ShouldBind(&calcReq); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

// calculates the math operation given
// uses a switch case internally.
// Returns an error if operator is not recognised.
func calculate(first float64, second float64, operator string) (float64, error) {
	switch operatingSystem {
	case "x":
		fmt.Println("Multiplication")
		// cases break automatically, no fallthrough by default
	case "/":
		fmt.Println("Division")
	case "+":
		fmt.Println("Addition")
	case "-":
		fmt.Println("Subtraction")
	default:
		// return error
		return 0, errNotSupportedOperator
	}
}
