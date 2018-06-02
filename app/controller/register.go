package controller

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	userLogout(w, r)
	if r.Method == "POST" {
		registerUser(w, r)
	} else {
		renderRegisterTemplate(w, r, false)
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	hashedPassword, _ := model.HashPassword(password)
	user := model.User{UserName: userName, Email: email, Password: hashedPassword}
	err := user.Add()
	if err != nil {
		renderRegisterTemplate(w, r, true)
		return
	}

	user, err = model.GetUserByName(userName)
	if err != nil {
		renderRegisterTemplate(w, r, true)
		return
	}
	session := model.CreateSession(user)
	cookie := &http.Cookie{
		Name:  "session",
		Value: session.ID,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/equipment", http.StatusSeeOther)
}

func renderRegisterTemplate(w http.ResponseWriter, r *http.Request, loginError bool) {
	t, err := template.New("register").ParseFiles("template/register.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)
	headerData := model.Header{Title: "Registrierung", Css: []string{"/css/style.css"}}
	footerData := []string{"scripts/validate.js"}
	registerData := model.RegisterData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: false,
		LoginError: loginError,
		FooterData: footerData}

	if err := t.Execute(w, registerData); err != nil {
		log.Fatalln(err)
	}
}
