package routes

import "io"
import "net/http"

func AdminsClients(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "AdminsClients")
}
