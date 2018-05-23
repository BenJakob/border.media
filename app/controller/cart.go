package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type CartData struct {
	HeaderData model.Header
	User       model.User
	IsLoggedIn bool
	Items      []model.Item
	FooterData []string
}

func Cart(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("cart").ParseFiles("template/cart.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Einkaufswagen", Css: []string{"/css/equipment.css", "/css/style.css"}}
	user := model.GetUser()
	items := model.GetEquipment()
	footerData := []string{}

	cartData := CartData{
		HeaderData: headerData,
		User:       user,
		IsLoggedIn: true,
		Items:      items,
		FooterData: footerData,
	}

	if err := t.Execute(w, cartData); err != nil {
		log.Fatalln(err)
	}
}
