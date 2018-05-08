package routes

import "io"
import "net/http"

func Cart(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Cart")
}
