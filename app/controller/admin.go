package routes

import "io"
import "net/http"

func Admin(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Admin")
}
