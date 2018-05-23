package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type AdminAddData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	FooterData []string
}

func AdminAdd(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-add").ParseFiles("template/admin-add.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Equipment hinzufügen", Css: []string{"/css/equipment.css", "/css/style.css"}}
	user := model.GetAdmin()
	footerData := []string{}

	adminAddData := AdminAddData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminAddData); err != nil {
		log.Fatalln(err)
	}
}
