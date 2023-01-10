// Package server test file
package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupContext() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// make sure c.Request is not nil
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	return c, w
}

func mockJSONPost(c *gin.Context, content interface{}) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonbytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	// the request body must be an io.ReadCloser
	// the bytes buffer though doesn't implement io.Closer,
	// so you wrap it in a no-op closer
	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func TestMultiplication(t *testing.T) {
	t.Run("Both positive", func(t *testing.T) {
		c, w := setupContext()
		// Set Body, Header and Content-Type
		mockJSONPost(c, &calcRequest{
			First:    2,
			Second:   3,
			Operator: "x",
		})

		// expected response
		expectedJSON := gin.H{
			"operation": "2.000000 x 3.000000",
			"result":    6,
			"status":    "success",
		}
		expected, _ := json.Marshal(expectedJSON)

		// call API endpoint
		calcEndpoint(c)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expected), w.Body.String())
	})

	t.Run("first negative", func(t *testing.T) {
		c, w := setupContext()
		// Set Body, Header and Content-Type
		mockJSONPost(c, &calcRequest{
			First:    -2,
			Second:   3,
			Operator: "x",
		})

		// expected response
		expectedJSON := gin.H{
			"operation": "-2.000000 x 3.000000",
			"result":    -6,
			"status":    "success",
		}
		expected, _ := json.Marshal(expectedJSON)

		// call API endpoint
		calcEndpoint(c)

		assert.Equal(t, 200, w.Code)
		assert.Equal(t, string(expected), w.Body.String())
	})
}
