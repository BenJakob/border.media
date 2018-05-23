package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type RegisterData struct {
	HeaderData model.Header
	IsLoggedIn bool
	FooterData []string
}

func Register(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("register").ParseFiles("template/register.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Registrierung", Css: []string{"/css/style.css"}}
	footerData := []string{}
	registerData := RegisterData{
		HeaderData: headerData,
		IsLoggedIn: false,
		FooterData: footerData}

	if err := t.Execute(w, registerData); err != nil {
		log.Fatalln(err)
	}
}
