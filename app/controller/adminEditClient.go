package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type AdminEditClientData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	Client     model.User
	FooterData []string
}

func AdminEditClient(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-edit-clients").ParseFiles("template/admin-edit-client.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Profil bearbeiten", Css: []string{"/css/profil.css", "/css/style.css"}}
	user := model.GetAdmin()
	client := model.GetUser()
	footerData := []string{}

	adminEditClientData := AdminEditClientData{
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
