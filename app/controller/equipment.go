package routes

import "io"
import "net/http"

func Equipment(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Equipment")
}
