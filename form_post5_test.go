package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	firstName := r.PostFormValue("firstName")
	middleName := r.PostFormValue("middleName")
	lastName := r.PostFormValue("lastName")

	fmt.Fprintf(w, "Hallo %s %s %s", firstName, middleName, lastName)
}

func TestPostForm(t *testing.T) {
	// untuk ngirim requestBody pake NewReader, artinya dengan atribut name firstname di tag html
	requestBody := strings.NewReader("firstName=Ilham&middleName=Muhammad&lastName=Sidiq")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	//headernya pake x-www-form-urlencoded karna ngirimnya lewat form
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	assert.Equal(t, "Hallo Ilham Muhammad Sidiq", string(body))
	assert.Equal(t, "application/x-www-form-urlencoded", request.Header.Get("content-type"))
}
