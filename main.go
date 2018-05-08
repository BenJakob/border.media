package main

import (
	"net/http"

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
	http.HandleFunc("/equipment", routes.Admin)
	http.HandleFunc("/login", routes.AdminAdd)
	http.HandleFunc("/my-equipment", routes.AdminsClients)
	http.HandleFunc("/profil", routes.AdminEditClients)
	http.HandleFunc("/register", routes.AdminEquipment)
	http.ListenAndServe(":8080", nil)

}
