package routes

import "io"
import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Index")
}
