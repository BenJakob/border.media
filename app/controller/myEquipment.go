package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"borgdir.media/app/model"
)

func MyEquipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("my-equipment").ParseFiles("template/my-equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	headerData := model.Header{Title: "Mein Equipment", Css: []string{"/css/style.css", "/css/equipment.css"}}
	rentedItems, markedItems := model.GetEntriesByCustomer(user)
	cartItemCount := model.GetCartItemCount(user)
	footerData := []string{"/scripts/my-equipment.js", "//cdnjs.cloudflare.com/ajax/libs/moment.js/2.7.0/moment.min.js"}

	myEquipmentData := model.MyEquipmentData{
		HeaderData:    headerData,
		User:          user,
		IsLoggedIn:    true,
		RentedItems:   rentedItems,
		MarkedItems:   markedItems,
		CartItemCount: cartItemCount,
		FooterData:    footerData,
	}

	if err := t.Execute(w, myEquipmentData); err != nil {
		log.Fatalln(err)
	}
}

func MyEquipmentRemove(w http.ResponseWriter, r *http.Request) {
	id := myEquipmentGetFormValue(w, r)

	if id != -1 {
		model.DeleteEntry(id)
	}

	http.Redirect(w, r, "/my-equipment", http.StatusSeeOther)
	return
}

func MyEquipmentExtend(w http.ResponseWriter, r *http.Request) {
	id := myEquipmentGetFormValue(w, r)

	if id != -1 {
		model.ExtendEntry(id)
		http.Redirect(w, r, "/my-equipment", http.StatusSeeOther)
	}
}

func myEquipmentGetFormValue(w http.ResponseWriter, r *http.Request) int {
	_, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return -1
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/my-equipment", http.StatusSeeOther)
		return -1
	}

	if !model.EntryExists(id) {
		http.Redirect(w, r, "/my-equipment", http.StatusSeeOther)
		return -1
	}

	return id
}
