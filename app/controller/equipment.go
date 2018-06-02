package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"borgdir.media/app/model"
)

func Equipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("equipment").ParseFiles("template/equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)
	isLoggedIn := err == nil
	headerData := model.Header{Title: "Meine Geräte", Css: []string{"/css/style.css", "/css/equipment.css"}}
	categories := model.GetCategories()
	items := model.GetItems()
	cartItemCount := model.GetCartItemCount(user)
	footerData := []string{"/scripts/equipment.js", "/scripts/search.js"}
	equipmentData := model.EquipmentData{
		HeaderData:    headerData,
		User:          user,
		IsLoggedIn:    isLoggedIn,
		Categories:    categories,
		Items:         items,
		CartItemCount: cartItemCount,
		FooterData:    footerData,
	}

	if err := t.Execute(w, equipmentData); err != nil {
		log.Fatalln(err)
	}
}

func EquipmentAdd(w http.ResponseWriter, r *http.Request) {
	user, item, err := getDataForEquipmentAction(w, r)
	if err == nil {
		model.CreateCartItem(user, item)
	}
	http.Redirect(w, r, "/equipment", http.StatusSeeOther)
}

func EquipmentMark(w http.ResponseWriter, r *http.Request) {
	user, item, err := getDataForEquipmentAction(w, r)
	if err == nil {
		date := time.Now().Format("20060102")
		model.CreateEntry(user, item, date, 2, 1)
	}
	http.Redirect(w, r, "/equipment", http.StatusSeeOther)
}

func getDataForEquipmentAction(w http.ResponseWriter, r *http.Request) (model.User, model.Item, error) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return model.User{}, model.Item{}, err
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return user, model.Item{}, err
	}

	item, err := model.GetItem(id)

	if err != nil {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return user, item, err
	}

	return user, item, err
}
