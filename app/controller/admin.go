package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type AdminData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	FooterData []string
}

func Admin(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin").ParseFiles("template/admin.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Verwaltung", Css: []string{"/css/style.css"}}
	user := model.GetAdmin()
	footerData := []string{}

	adminData := AdminData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminData); err != nil {
		log.Fatalln(err)
	}
}
