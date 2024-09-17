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

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	//bisa inject data pake map atau struct
	err := t.ExecuteTemplate(w, "simple.gohtml", map[string]interface{}{
		"Title": "Hello World",
		"Name":  "Ilham",
	})
	if err != nil {
		return
	}
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}

type Page struct {
	Title string
	Name  string
	Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	err := t.ExecuteTemplate(w, "simple.gohtml", Page{
		Title: "Hello Struct",
		Name:  "Joko",
		Address: Address{
			Street: "Kp Setu",
		},
	})
	if err != nil {
		return
	}
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}
