package controller

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

func AdminsClients(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-clients").ParseFiles("template/admin-clients.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")
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

	headerData := model.Header{Title: "Kunden verwalten", Css: []string{"/css/equipment.css", "/css/style.css"}}
	users := model.GetUsers()
	userStatus := model.GetStatus()
	footerData := []string{"/scripts/search.js"}

	adminClientsData := model.AdminClientsData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Users:      users,
		UserStatus: userStatus,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminClientsData); err != nil {
		log.Fatalln(err)
	}
}
