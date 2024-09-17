package restful_api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Title string
	Name  string
}

func (m MyPage) SayHello(name string) string {
	return fmt.Sprintf("Hello %s my name is %s", name, m.Name)
}

func TemplateFunction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/function.gohtml"))
	err := t.ExecuteTemplate(w, "function.gohtml", MyPage{
		Title: "Template Function",
		Name:  "Ilham",
	})
	if err != nil {
		return
	}
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}

type Mhs struct {
	Name string
}

func FunctionCreateGlobal(w http.ResponseWriter, r *http.Request) {
	t := template.New("FUNCTION") // bikin nama templaenya
	t = t.Funcs(map[string]any{   // bikin functionnya
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	//jadi ini adalah function pipeline, jadi hasil dri function sayHello, masuk ke parameter function upper
	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`)) // parse templatenya
	err := t.ExecuteTemplate(w, "FUNCTION", Mhs{             // execute templatenya
		Name: "hamzah",
	})
	if err != nil {
		return
	}
}

func TestFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	FunctionCreateGlobal(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)

	assert.Nil(t, err)
	fmt.Println(string(body))
}
