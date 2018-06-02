package controller

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin").ParseFiles("template/admin.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

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

	headerData := model.Header{Title: "Verwaltung", Css: []string{"/css/style.css"}}
	footerData := []string{}
	adminData := model.AdminData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminData); err != nil {
		log.Fatalln(err)
	}
}
