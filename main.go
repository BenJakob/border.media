package main

import (
	"bufio"
	"net/http"
	"os"
	"strings"

	"borgdir.media/app/controller"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/admin", controller.Admin)
	http.HandleFunc("/admin/add", controller.AdminAdd)
	http.HandleFunc("/admin/clients", controller.AdminsClients)
	http.HandleFunc("/admin/edit-client", controller.AdminEditClient)
	http.HandleFunc("/admin/edit-client/lock", controller.AdminEditClientLock)
	http.HandleFunc("/admin/equipment", controller.AdminEquipment)
	http.HandleFunc("/admin/equipment/edit", controller.AdminEquipmentEdit)
	http.HandleFunc("/admin/equipment/delete", controller.AdminEquipmentDelete)
	http.HandleFunc("/cart", controller.Cart)
	http.HandleFunc("/cart/remove", controller.CartRemove)
	http.HandleFunc("/cart/updatedate", controller.CartUpdateDate)
	http.HandleFunc("/cart/updatequantity", controller.CartUpdateQuantity)
	http.HandleFunc("/cart/checkout", controller.CartCheckout)
	http.HandleFunc("/equipment", controller.Equipment)
	http.HandleFunc("/equipment/add", controller.EquipmentAdd)
	http.HandleFunc("/equipment/mark", controller.EquipmentMark)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/my-equipment", controller.MyEquipment)
	http.HandleFunc("/my-equipment/remove", controller.MyEquipmentRemove)
	http.HandleFunc("/my-equipment/extend", controller.MyEquipmentExtend)
	http.HandleFunc("/profil", controller.Profil)
	http.HandleFunc("/profil/delete", controller.ProfilDelete)
	http.HandleFunc("/register", controller.Register)

	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/scripts/", serveResource)

	http.ListenAndServe(":8080", nil)
}

func serveResource(w http.ResponseWriter, r *http.Request) {
	path := "static" + r.URL.Path
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
