package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type ProfilData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	FooterData []string
}

func Profil(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("profil").ParseFiles("template/profil.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Profil", Css: []string{"/css/profil.css", "/css/style.css"}}
	user := model.GetUser()
	footerData := []string{}

	profilData := ProfilData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		FooterData: footerData,
	}

	if err := t.Execute(w, profilData); err != nil {
		log.Fatalln(err)
	}
}
