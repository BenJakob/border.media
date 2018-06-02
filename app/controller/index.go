package controller

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("content").ParseFiles("template/index.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	user, err := model.GetLoggedInUser(w, r)
	var isLoggedIn bool
	if err != nil {
		isLoggedIn = false
	} else {
		isLoggedIn = true
	}

	headerData := model.Header{Title: "borgdir.media", Css: []string{"/css/index.css", "/css/style.css"}}
	activeImg, images := model.GetCarouselImages()
	cartItemCount := model.GetCartItemCount(user)
	footerData := []string{"scripts/index.js"}
	indexData := model.IndexData{
		HeaderData:    headerData,
		User:          user,
		IsLoggedIn:    isLoggedIn,
		ActiveImg:     activeImg,
		CartItemCount: cartItemCount,
		Images:        images,
		FooterData:    footerData,
	}

	if err := t.Execute(w, indexData); err != nil {
		log.Fatalln(err)
	}
}
