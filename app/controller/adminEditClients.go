package routes

import "io"
import "net/http"

func AdminEditClients(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "AdminEditClients")
}
