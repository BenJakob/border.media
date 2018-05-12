package routes

import (
	"html/template"
	"net/http"

	"borgdir.media/app/model"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	// t := template.Must(template.ParseFiles("template/admin.html"))

	t := template.Must(template.ParseFiles("template/layout.gohtml"))
	headerData := model.Header{Title: "Admin", Css: []string{}}
	navbarItems := model.NavItem{Title: "Equipment", Link: "admin-equipment.html"}

	t.Execute(w, headerData)

}
