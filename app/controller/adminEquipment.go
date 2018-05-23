package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type AdminEquipmentData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	Items      []model.Item
	FooterData []string
}

func AdminEquipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("admin-equipment").ParseFiles("template/admin-equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Equipment verwalten", Css: []string{"/css/equipment.css", "/css/style.css"}}
	user := model.GetAdmin()
	items := model.GetEquipment()
	footerData := []string{}

	adminEquipmentData := AdminEquipmentData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Items:      items,
		FooterData: footerData,
	}

	if err := t.Execute(w, adminEquipmentData); err != nil {
		log.Fatalln(err)
	}
}
