package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type MyEquipmentData struct {
	HeaderData  model.Header
	User        model.User
	IsLoggedIn  bool
	RentedItems []model.Item
	MarkedItems []model.Item
	FooterData  []string
}

func MyEquipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("my-equipment").ParseFiles("template/my-equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Mein Equipment", Css: []string{"/css/equipment.css", "/css/style.css"}}
	user := model.GetUser()
	rentedItems := model.GetEquipment()
	markedItems := model.GetEquipment()
	footerData := []string{}

	myEquipmentData := MyEquipmentData{
		HeaderData:  headerData,
		User:        user,
		IsLoggedIn:  true,
		RentedItems: rentedItems,
		MarkedItems: markedItems,
		FooterData:  footerData,
	}

	if err := t.Execute(w, myEquipmentData); err != nil {
		log.Fatalln(err)
	}
}
