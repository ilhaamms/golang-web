package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type") // ambil header dari client
	fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/login", nil)
	request.Header.Add("content-type", "application/json") // mengirim request ke server dengan menambahkan header

	response := httptest.NewRecorder()

	RequestHeader(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, "application/json", string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	// custom header dari server ke client, atau bisa custom sebaliknya dari client ke server
	w.Header().Add("X-Powered-By", "Ilham Muhammad Sidiq")
	fmt.Fprintln(w, "OK")
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	response := httptest.NewRecorder()

	ResponseHeader(response, request)

	result := response.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, "OK\n", string(body))
	assert.Equal(t, "Ilham Muhammad Sidiq", response.Header().Get("X-Powered-By"))
}
