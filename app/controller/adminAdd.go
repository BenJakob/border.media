package routes

import "io"
import "net/http"

func AdminAdd(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Admin")
}
