package routes

import "io"
import "net/http"

func Profil(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Profil")
}
