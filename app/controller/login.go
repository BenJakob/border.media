package controller

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if isLoggedIn(w, r) {
		http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		return
	}

	if r.Method == "POST" {
		userLogin(w, r)
	} else {
		renderTemplate(w, r, false)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	userLogout(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func isLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("session")

	if err != nil {
		return false
	}

	if !model.SessionExists(cookie.Value) {
		userLogout(w, r)
		return false
	}

	return true
}

func userLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	password := r.FormValue("password")
	user, err := model.GetUserByName(userName)
	if err != nil {
		renderTemplate(w, r, true)
		return
	}

	if model.CheckPasswordHash(password, user.Password) {
		session := model.CreateSession(user)
		cookie := &http.Cookie{
			Name:  "session",
			Value: session.ID,
		}
		http.SetCookie(w, cookie)
		if user.Status == "Benutzer" {
			http.Redirect(w, r, "/equipment", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		}
		return
	}

	renderTemplate(w, r, true)
}

func userLogout(w http.ResponseWriter, r *http.Request) {
	currCookie, err := r.Cookie("session")

	if err != nil {
		return
	}

	model.DeleteSession(currCookie.Value)

	delCookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, delCookie)
}

func renderTemplate(w http.ResponseWriter, r *http.Request, loginError bool) {
	t, err := template.New("login").ParseFiles("template/login.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Login", Css: []string{"/css/style.css"}}
	footerData := []string{}
	loginData := model.LoginData{
		HeaderData: headerData,
		IsLoggedIn: false,
		LoginError: loginError,
		FooterData: footerData,
	}

	if err := t.Execute(w, loginData); err != nil {
		log.Fatalln(err)
	}
}
