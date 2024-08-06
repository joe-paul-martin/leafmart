package handlers

import (
	"net/http"
	"time"
)

func HomePage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to homepage\n"))

	time.Sleep(1 * time.Second)

	w.Write([]byte("Starting the go routine\n"))

	ctx := req.Context()

	for {
		select {
		case <-ctx.Done():
			w.Write([]byte("Stopping the response writer as 2 second has been reached\n"))
			return
		default:
			// fmt.Println("Continuing writing since context deadline has not yet reached\n")
		}
	}
}

func SecondPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome to secondPage\n"))

	ctx := req.Context()

	value := ctx.Value("id").(string)

	time.Sleep(1 * time.Second)

	w.Write([]byte(value))

}
