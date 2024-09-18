package restful_api

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct {
	Handler http.Handler
}

func (e ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		err := recover()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error %s", err)
		}
	}()

	e.Handler.ServeHTTP(writer, request)
}

func (l *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before execute handler")
	l.Handler.ServeHTTP(writer, request)
	fmt.Println("After execute handler")
}

func TestMiddleware(t *testing.T) {
	/*
		intinya sistem middleware cara kerjanya yaitu
		middleware -> handler -> balik lagi ke middleware
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Execute Handler")
	})

	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		panic("UPS")
	})

	//karna LogMiddleware mengimplementasikan ServeHTTP, maka kita wajib pointer receiver dengan seperti ini
	//yaitu tanda &
	logMiddleware := &LogMiddleware{
		Handler: mux, // dan keempat log middleware akan ngirim request ke mux
	}
	// nah kalau urutannya dari middleware yang akhir yaitu
	// logMiddleware -> errorHandler

	errorHandler := &ErrorHandler{
		Handler: logMiddleware, // ketiga ke log middleware
	}

	//urutan request yang masuk yaitu
	server := http.Server{ // pertama ke server, lalu server ngirim request ke
		Addr:    "localhost:8080",
		Handler: errorHandler, // kedua ke errorHandler, kemudian errorHandler akan ngrim request ke logMiddleware
	}

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
