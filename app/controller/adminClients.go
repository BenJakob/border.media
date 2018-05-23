package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type AdminClientsData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	Users      []model.User
	FooterData []string
}

func AdminsClients(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-clients").ParseFiles("template/admin-clients.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Kunden verwalten", Css: []string{"/css/equipment.css", "/css/style.css"}}
	user := model.GetAdmin()
	users := model.GetUsers()
	footerData := []string{}

	adminClientsData := AdminClientsData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Users:      users,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminClientsData); err != nil {
		log.Fatalln(err)
	}
}
