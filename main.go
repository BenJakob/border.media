package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"borgdir.media/app/controller"
)

func main() {

	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/admin", routes.Admin)
	http.HandleFunc("/admin/add", routes.AdminAdd)
	http.HandleFunc("/admin/clients", routes.AdminsClients)
	http.HandleFunc("/admin/edit-client", routes.AdminEditClients)
	http.HandleFunc("/admin/equipment", routes.AdminEquipment)
	http.HandleFunc("/cart", routes.Index)
	http.HandleFunc("/equipment", routes.Equipment)
	http.HandleFunc("/login", routes.Login)
	http.HandleFunc("/my-equipment", routes.MyEquipment)
	http.HandleFunc("/profil", routes.Profil)
	http.HandleFunc("/register", routes.Register)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/scripts/", serveResource)

	http.ListenAndServe(":8080", nil)

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	path := "static" + req.URL.Path
	var contentType string

	if strings.HasSuffix(path, ".css") {
		contentType = "text/css; charset=utf-8"
	} else if strings.HasSuffix(path, ".js") {
		contentType = "application/javascript; charset=utf-8"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png; charset=utf-8"
	} else if strings.HasSuffix(path, ".jpg") {
		contentType = "image/jpg; charset=utf-8"
	} else if strings.HasSuffix(path, ".svg") {
		contentType = "image/svg+xml; charset=utf-8"
	} else {
		contentType = "text/plain; charset=utf-8"
	}

	f, err := os.Open(path)
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}
