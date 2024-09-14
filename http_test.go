package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Response")
}

func TestHttp(t *testing.T) {
	// ini untuk request di unit testing
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/wkwk", nil)
	// ini untuk response di unit testing
	response := httptest.NewRecorder()

	//setelah di ser request dan responsenya, jangan lupa panggil method handlernya
	HelloHandler(response, request)

	result := response.Result()        // baca response
	body, _ := io.ReadAll(result.Body) // baca body
	dataString := string(body)

	assert.Equal(t, "Hello Response", dataString)
	require.Equal(t, 200, response.Code)
}
