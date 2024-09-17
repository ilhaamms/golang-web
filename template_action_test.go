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

type AddressWith struct {
	Street string
	City   string
}

type PageWith struct {
	Title string
	Name  string
	AddressWith
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	//kalo nested struct pake with di template gohtml nya
	t.ExecuteTemplate(w, "address.gohtml", PageWith{
		Title: "address template",
		Name:  "Ilham Muhammad Sidiq",
		AddressWith: AddressWith{
			Street: "Kp Setu",
			City:   "Bekasi",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}
