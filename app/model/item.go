package model

import (
	"time"
)

type Item struct {
	ID          uint32
	Name        string
	Description string
	Image       string
	Location    string
	Category    string
	Quantity    int
	Available   bool
	Price       float64
	Orders      []LedgerEntry
}

func GetCarouselImages() (string, []string) {
	rows, err := db.Query("SELECT image FROM item LIMIT 6")
	checkErr(err)

	var images []string

	for rows.Next() {
		var imagePath string
		err = rows.Scan(&imagePath)
		checkErr(err)
		images = append(images, imagePath)
	}

	if len(images) > 0 {
		activeImage := images[0]
		images = append(images[:0], images[1:]...)
		return activeImage, images
	}

	return "", images
}

func GetCategories() []string {
	rows, err := db.Query("SELECT Name FROM [Item Category]")
	checkErr(err)

	var categories []string

	for rows.Next() {
		var category string
		err = rows.Scan(&category)
		checkErr(err)
		categories = append(categories, category)
	}
	rows.Close()
	return categories
}

func GetLocations() []string {
	rows, err := db.Query("SELECT Name FROM Location")
	checkErr(err)

	var locations []string

	for rows.Next() {
		var location string
		err = rows.Scan(&location)
		checkErr(err)
		locations = append(locations, location)
	}
	rows.Close()
	return locations
}

func GetItem(id int) (Item, error) {
	item := Item{}
	timeNow := time.Now().Format(dbDateLayout)
	var quantityAvailable int
	err := db.QueryRow(`
		SELECT Item.ID, Item.Name, Item.Description, Item.Image, Location.Name as Location, [Item Category].Name AS Category, Item.Price, Item.Quantity, Item.Quantity-IFNULL(SUM(Stock.Quantity), 0) AS [Quantity Available]
		FROM item
		INNER JOIN [Item Category] ON Item.Category = [Item Category].ID
		INNER JOIN Location ON Item.Location = Location.ID
		LEFT JOIN [Ledger Entry] AS Stock on Item.ID = Stock.[Item ID] AND (Stock.Type = 2 OR Stock.Type = 3) AND [Start Date] >= $1 AND [End Date] <= $2
		WHERE Item.ID = $3
		GROUP BY item.id`, timeNow, timeNow, id).Scan(&item.ID, &item.Name, &item.Description, &item.Image, &item.Location, &item.Category, &item.Price, &item.Quantity, &quantityAvailable)
	checkErr(err)

	if quantityAvailable > 0 {
		item.Available = true
	} else {
		item.Available = false
	}

	return item, err
}

func GetItems() []Item {
	timeNow := time.Now().Format(dbDateLayout)
	var quantityAvailable int
	rows, err := db.Query(`
		SELECT Item.ID, Item.Name, Item.Description, Item.Image, Location.Name as Location, [Item Category].Name AS Category, Item.Price, Item.Quantity, Item.Quantity-IFNULL(SUM(Stock.Quantity), 0) AS [Quantity Available]
		FROM item
		INNER JOIN [Item Category] ON Item.Category = [Item Category].ID
		INNER JOIN Location ON Item.Location = Location.ID
		LEFT JOIN [Ledger Entry] AS Stock on Item.ID = Stock.[Item ID] AND (Stock.Type = 2 OR Stock.Type = 3) AND [Start Date] >= $1 AND [End Date] <= $2
		GROUP BY item.id`, timeNow, timeNow,
	)
	checkErr(err)

	var items []Item

	for rows.Next() {
		var item = Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Image, &item.Location, &item.Category, &item.Price, &item.Quantity, &quantityAvailable)
		checkErr(err)
		if quantityAvailable > 0 {
			item.Available = true
		} else {
			item.Available = false
		}
		items = append(items, item)
	}

	return items
}

func GetItemsAndOrders() []Item {
	timeNow := time.Now().Format(dbDateLayout)
	var quantityAvailable int
	rows, err := db.Query(`
		SELECT Item.ID, Item.Name, Item.Description, Item.Image, Location.Name as Location, [Item Category].Name AS Category, Item.Price, Item.Quantity, Item.Quantity-IFNULL(SUM(Stock.Quantity), 0) AS [Quantity Available]
		FROM item
		INNER JOIN [Item Category] ON Item.Category = [Item Category].ID
		INNER JOIN Location ON Item.Location = Location.ID
		LEFT JOIN [Ledger Entry] AS Stock on Item.ID = Stock.[Item ID] AND (Stock.Type = 2 OR Stock.Type = 3) AND [Start Date] >= $1 AND [End Date] <= $2
		GROUP BY item.id`, timeNow, timeNow,
	)
	checkErr(err)

	var items []Item

	for rows.Next() {
		var item = Item{}
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Image, &item.Location, &item.Category, &item.Price, &item.Quantity, &quantityAvailable)
		checkErr(err)
		if quantityAvailable > 0 {
			item.Available = true
		} else {
			item.Available = false
		}
		item.Orders = getEntriesByItem(item, 3)
		items = append(items, item)
	}

	return items
}

func (item *Item) Update() error {
	var location int
	var category int
	err := db.QueryRow(`
		SELECT ID
		FROM Location
		WHERE Name = $1`, item.Location).Scan(&location)
	checkErr(err)
	err = db.QueryRow(`
		SELECT ID
		FROM [Item Category]
		WHERE Name = $1`, item.Category).Scan(&category)
	checkErr(err)
	_, err = db.Exec(`
		UPDATE Item SET Name=$1, Description=$2, Location=$3, Category=$4, Image=$5, Price=$6, Quantity=$7
		WHERE ID = $8`, item.Name, item.Description, location, category, item.Image, item.Price, item.Quantity, item.ID)
	return err
}

func (item *Item) Add() error {
	var location int
	var category int
	err := db.QueryRow(`
		SELECT ID
		FROM Location
		WHERE Name = $1`, item.Location).Scan(&location)
	checkErr(err)
	err = db.QueryRow(`
		SELECT ID
		FROM [Item Category]
		WHERE Name = $1`, item.Category).Scan(&category)
	checkErr(err)
	_, err = db.Exec(`
		INSERT INTO Item(Name, Description, Image, Location, Category, Price, Quantity)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		`, item.Name, item.Description, item.Image, location, category, item.Price, item.Quantity)
	return err
}

func (item *Item) Delete() {
	_, err = db.Exec("DELETE FROM Item WHERE ID = $1", item.ID)
	checkErr(err)
}
