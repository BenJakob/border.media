package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type LoginData struct {
	HeaderData model.Header
	IsLoggedIn bool
	FooterData []string
}

func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("login").ParseFiles("template/login.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Login", Css: []string{"/css/style.css"}}
	footerData := []string{}
	indexData := IndexData{
		HeaderData: headerData,
		IsLoggedIn: false,
		FooterData: footerData,
	}

	if err := t.Execute(w, indexData); err != nil {
		log.Fatalln(err)
	}
}
