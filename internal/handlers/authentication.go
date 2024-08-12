package handlers

import "net/http"

func AuthenticationPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("\n from the authentication page"))
}
