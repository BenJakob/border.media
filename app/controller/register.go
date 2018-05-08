package routes

import "io"
import "net/http"

func Register(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Register")
}
