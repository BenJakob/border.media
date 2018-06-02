package controller

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

func Profil(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		profilUpdate(w, r, user)
	} else {
		renderProfilTemplate(w, r, user, false)
	}
}

func profilUpdate(w http.ResponseWriter, r *http.Request, user model.User) {
	userName := r.FormValue("username")
	email := r.FormValue("email")
	oldPassword := r.FormValue("old-password")
	newPassword := r.FormValue("password")
	hashedPassword, err := model.HashPassword(newPassword)

	if err != nil {
		renderProfilTemplate(w, r, user, true)
		return
	}

	if oldPassword != "" {
		if model.CheckPasswordHash(oldPassword, user.Password) {
			user.Password = hashedPassword
		} else {
			renderProfilTemplate(w, r, user, true)
			return
		}
	}

	user.UserName = userName
	user.Email = email
	err = user.Update()
	if err != nil {
		renderProfilTemplate(w, r, user, true)
		return
	}

	renderProfilTemplate(w, r, user, false)
}

func renderProfilTemplate(w http.ResponseWriter, r *http.Request, user model.User, loginError bool) {
	t, err := template.New("profil").ParseFiles("template/profil.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Profil", Css: []string{"/css/style.css"}}
	cartItemCount := model.GetCartItemCount(user)
	footerData := []string{"/scripts/validate.js", "/scripts/profil.js"}

	profilData := model.ProfilData{
		HeaderData:    headerData,
		User:          user,
		IsLoggedIn:    true,
		LoginError:    loginError,
		CartItemCount: cartItemCount,
		FooterData:    footerData,
	}

	if err := t.Execute(w, profilData); err != nil {
		log.Fatalln(err)
	}
}

func ProfilDelete(w http.ResponseWriter, r *http.Request) {
	user, err := model.GetLoggedInUser(w, r)

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	Logout(w, r)
	user.Delete()
}
