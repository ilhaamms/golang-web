package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "silahkan masukan query parameter")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Selamat datang %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	response := httptest.NewRecorder()

	ResponseCode(response, request)

	result := response.Result()
	body, err := io.ReadAll(result.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, result.StatusCode)
	assert.Equal(t, "silahkan masukan query parameter", string(body))
}

func TestResponseCodeSuccess(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=Ilham", nil)
	response := httptest.NewRecorder()

	ResponseCode(response, request)

	result := response.Result()
	body, err := io.ReadAll(result.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, "Selamat datang Ilham", string(body))
}
