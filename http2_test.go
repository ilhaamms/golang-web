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

	// ini untuk "rekaman" virtual yang menyimpan apa yang ditulis oleh handler(server) ke respons,
	// sehingga Anda bisa memeriksanya di dalam unit test.
	recorder := httptest.NewRecorder()

	//setelah di set request, jangan lupa panggil method handlernya
	HelloHandler(recorder, request)

	response := recorder.Result()        // baca response
	body, _ := io.ReadAll(response.Body) // baca response body
	dataString := string(body)

	assert.Equal(t, "Hello Response", dataString)
	require.Equal(t, 200, response.StatusCode) // defaultnya kalau status code ga di set adalah 200
}
