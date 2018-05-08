package routes

import "io"
import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Login")
}
