package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	//kalau banyak html untuk import template maka wajib disebut semuua templatenya
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
	))

	//kalo nested struct pake with di template gohtml nya
	t.ExecuteTemplate(w, "layout", PageWith{
		Title: "layout template",
		Name:  "Hamzah Muhammad Ramadhan",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}
