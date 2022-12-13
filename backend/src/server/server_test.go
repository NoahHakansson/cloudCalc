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

func TestSomething(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	assert.Nil(t, nil)

	// assert for not nil (good when you expect something)
	if assert.NotNil(t, nil) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", "Something")
	}
}

func TestMultiplication(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// make sure c.Request is not nil
	c.Request = &http.Request{
		Header: make(http.Header),
	}

	// Set Body, Header and Content-Type
	mockJSONPost(c, &calcRequest{
		First:    2,
		Second:   3,
		Operator: "x",
	})

	// expected response
	expectedJSON := gin.H{
		"result": 6,
		"status": "success",
	}
	expected, _ := json.Marshal(expectedJSON)

	// call API endpoint
	calcEndpoint(c)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expected), w.Body.String())
}
