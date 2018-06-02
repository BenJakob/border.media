package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"borgdir.media/app/model"
)

func AdminAdd(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if user.Status != "Admin" {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		addItem(w, r, user)
	} else {
		renderAdminAddTemplate(w, r, user)
	}
}

func addItem(w http.ResponseWriter, r *http.Request, user model.User) {
	name := r.FormValue("name")
	description := r.FormValue("description")
	category := r.FormValue("category")
	location := r.FormValue("location")
	price, err := strconv.ParseFloat(r.FormValue("price"), 64)
	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}

	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}

	item := model.Item{
		Name:        name,
		Description: description,
		Category:    category,
		Location:    location,
		Price:       price,
		Quantity:    quantity,
		Image:       "/img/image.svg",
	}

	err = item.Add()
	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
}

func renderAdminAddTemplate(w http.ResponseWriter, r *http.Request, user model.User) {
	t, err := template.New("admin-add").ParseFiles("template/admin-add.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Equipment hinzufügen", Css: []string{"/css/equipment.css", "/css/style.css"}}
	categories := model.GetCategories()
	locations := model.GetLocations()
	footerData := []string{}
	adminAddData := model.AdminAddData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Categories: categories,
		Locations:  locations,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminAddData); err != nil {
		log.Fatalln(err)
	}
}
