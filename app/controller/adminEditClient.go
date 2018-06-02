package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"borgdir.media/app/model"
)

func AdminEditClient(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-edit-clients").ParseFiles("template/admin-edit-client.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if user.Status != "Admin" {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/admin/clients", http.StatusSeeOther)
		return
	}

	client, err := model.GetUser(uint32(id))

	if err != nil {
		http.Redirect(w, r, "/admin/clients", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		client = adminEditClientUpdate(r, client)
	}

	headerData := model.Header{Title: "Profil bearbeiten", Css: []string{"/css/style.css"}}
	footerData := []string{"/scripts/admin-edit-client.js", "/scripts/validate.js"}

	adminEditClientData := model.AdminEditClientData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Client:     client,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminEditClientData); err != nil {
		log.Fatalln(err)
	}
}

func AdminEditClientLock(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

	if user.Status != "Admin" {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Redirect(w, r, "/admin/clients", http.StatusSeeOther)
		return
	}

	client, err := model.GetUser(uint32(id))

	if err != nil {
		http.Redirect(w, r, "/admin/clients", http.StatusSeeOther)
		return
	}

	client.Status = "Gesperrt"

	client.Update()
}

func adminEditClientUpdate(r *http.Request, client model.User) model.User {
	userName := r.FormValue("username")
	email := r.FormValue("email")
	newPassword := r.FormValue("password")

	if newPassword != "" {
		hashedPassword, _ := model.HashPassword(newPassword)
		client.Password = hashedPassword
	}

	client.UserName = userName
	client.Email = email
	client.Update()
	return client
}
