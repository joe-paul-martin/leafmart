package handlers

import (
	"net/http"
	"time"
)

func HomePage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to homepage\n"))

}

func SecondPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to secondPage\n"))

	ctx := req.Context()

	value := ctx.Value("id").(string)

	time.Sleep(1 * time.Second)

	w.Write([]byte(value))

}

func AdminPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to admin Page\n"))

}
