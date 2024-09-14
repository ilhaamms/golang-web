package restful_api

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	//HandlerFunc ini gabisa bikin banyak endpoint URL, jadi cuman di nama domain aja 1
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func TestServeMux(t *testing.T) {
	// note NewServeMux sama aja dengan ServeMux, tapi paling sering dipake yg NewServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Selamat Datang")
	})

	mux.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Silahkan login")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
