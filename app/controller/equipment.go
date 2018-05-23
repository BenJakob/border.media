package routes

import (
	"html/template"
	"log"
	"net/http"

	"borgdir.media/app/model"
)

type EquipmentData struct {
	HeaderData  model.Header
	User        model.User
	IsLoggedIn  bool
	Categories  []string
	SortOptions []string
	Rows        []Row
	FooterData  []string
}

type Row struct {
	Items []model.Item
}

func Equipment(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("equipment").ParseFiles("template/equipment.gohtml", "template/header.gohtml", "template/navbar.gohtml", "template/footer.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	headerData := model.Header{Title: "Meine Geräte", Css: []string{"/css/equipment.css", "/css/style.css"}}
	user := model.GetUser()
	categories := model.GetCategories()
	sortOptions := model.GetSortOptions()
	rows := createRows(model.GetEquipment())
	footerData := []string{}

	equipmentData := EquipmentData{
		HeaderData:  headerData,
		User:        user,
		IsLoggedIn:  true,
		Categories:  categories,
		SortOptions: sortOptions,
		Rows:        rows,
		FooterData:  footerData,
	}

	if err := t.Execute(w, equipmentData); err != nil {
		log.Fatalln(err)
	}
}

func createRows(items []model.Item) []Row {
	var rows []Row
	var row Row
	for index, item := range items {
		if index%2 == 0 {
			row = Row{}
		}
		row.Items = append(row.Items, item)
		if index%2 == 1 || index+1 == len(items) {
			rows = append(rows, row)
		}
	}
	return rows
}
