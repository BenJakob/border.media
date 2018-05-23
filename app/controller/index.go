package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type IndexData struct {
	HeaderData model.Header
	IsLoggedIn bool
	ActiveImg  string
	Images     []string
	FooterData []string
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("content").ParseFiles("template/index.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "borgdir.media", Css: []string{"/css/index.css", "/css/style.css"}}
	activeImg, images := model.GetCarouselImages()
	footerData := []string{"scripts/index.js"}
	indexData := IndexData{
		HeaderData: headerData,
		IsLoggedIn: false,
		ActiveImg:  activeImg,
		Images:     images,
		FooterData: footerData,
	}

	if err := t.Execute(w, indexData); err != nil {
		log.Fatalln(err)
	}
}
