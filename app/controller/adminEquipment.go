package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"borgdir.media/app/model"
)

func AdminEquipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-equipment").ParseFiles("template/admin-equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if user.Status != "Admin" {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return
	}

	headerData := model.Header{Title: "Equipment verwalten", Css: []string{"/css/style.css", "/css/equipment.css"}}
	categories := model.GetCategories()
	items := model.GetItemsAndOrders()
	footerData := []string{"/scripts/admin-equipment.js", "/scripts/search.js"}

	adminEquipmentData := model.AdminEquipmentData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Categories: categories,
		Items:      items,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminEquipmentData); err != nil {
		log.Fatalln(err)
	}
}

func AdminEquipmentDelete(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if user.Status != "Admin" {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}

	item, err := model.GetItem(id)
	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}

	item.Delete()
	http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
}

func AdminEquipmentEdit(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if user.Status != "Admin" {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}

	item, err := model.GetItem(id)
	if err != nil {
		http.Redirect(w, r, "/admin/equipment", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
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

		item.Name = name
		item.Description = description
		item.Category = category
		item.Location = location
		item.Price = price
		item.Quantity = quantity

		item.Update()
	}
	renderEquipmentEditTemplate(w, r, item, user)
}

func renderEquipmentEditTemplate(w http.ResponseWriter, r *http.Request, item model.Item, user model.User) {
	t, err := template.New("admin-edit-equipment").ParseFiles("template/admin-edit-equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Equipment bearbeiten", Css: []string{"/css/style.css", "/css/equipment.css"}}
	categories := remove(model.GetCategories(), item.Category)
	categories = insertAt(0, categories, item.Category)
	locations := remove(model.GetLocations(), item.Location)
	locations = insertAt(0, locations, item.Location)
	footerData := []string{"/scripts/admin-edit-equipment.js"}
	adminEditEquipmentData := model.AdminEditEquipmentData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Item:       item,
		Categories: categories,
		Locations:  locations,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminEditEquipmentData); err != nil {
		log.Fatalln(err)
	}
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func insertAt(index int, s []string, r string) []string {
	s = append(s, "")
	copy(s[index+1:], s[index:])
	s[index] = r
	return s
}
