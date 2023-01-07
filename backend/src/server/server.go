// Package server provides server
package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// TODO: move calculator stuff to its own package

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
func Start() (*gin.Engine, error) {
	r := gin.Default()
	// set trusted proxy to localhost
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	// routes
	r.Use(cors.Default())
	r.POST("/api/calc", calcEndpoint)

	return r, nil
}

func calcEndpoint(c *gin.Context) {
	calcReq := calcRequest{}
	// bind body data or return error if it fails
	if err := c.ShouldBind(&calcReq); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// calculate result
	result, err := calculate(calcReq.First, calcReq.Second, calcReq.Operator)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":    err.Error(),
			"first":    calcReq.First,
			"second":   calcReq.Second,
			"operator": calcReq.Operator,
		})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*")
	// return result
	println(
		"First: ", int(calcReq.First),
		"Second: ", int(calcReq.Second),
		"Operator: ", calcReq.Operator,
		"Result: ", int(result),
	)
	c.JSON(http.StatusOK, gin.H{
		"status":    "success",
		"operation": fmt.Sprintf("%f %s %f", calcReq.First, calcReq.Operator, calcReq.Second),
		"result":    result,
	})
}

// calculates the math operation given
// uses a switch case internally.
// Returns an error if operator is not recognised.
func calculate(first float64, second float64, operator string) (float64, error) {
	var result float64
	switch operator {
	// cases break automatically, no fallthrough by default
	case "^":
		fmt.Println("Subtraction")
    if second != 0 {
      result = first
      for i := 0.0;i < second;i++ {
        result = result * first
      }
    } else {
      result = 1
    }
	case "x":
		fmt.Println("Multiplication")
		result = first * second
	case "/":
		fmt.Println("Division")
		result = first / second
	case "+":
		fmt.Println("Addition")
		result = first + second
	case "-":
		fmt.Println("Subtraction")
		result = first - second
	default:
		return 0, errNotSupportedOperator
	}
	return result, nil
}
