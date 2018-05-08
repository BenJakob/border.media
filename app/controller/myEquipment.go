package routes

import "io"
import "net/http"

func MyEquipment(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "MyEquipment")
}
