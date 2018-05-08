package routes

import "io"
import "net/http"

func AdminEquipment(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "AdminEquipment")
}
