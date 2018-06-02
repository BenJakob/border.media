package model

import "time"

type CartItem struct {
	ID          int
	UserID      int
	ItemID      int
	Name        string
	Description string
	Image       string
	Quantity    int
	EndDate     string
}

func CreateCartItem(user User, item Item) {
	statement := "INSERT INTO Cart([User ID], [Item ID], Quantity, [End Date]) VALUES ($1, $2, $3, $4)"
	stmt, err := db.Prepare(statement)
	checkErr(err)

	endDate := time.Now().AddDate(0, 0, 7).Format(dbDateLayout)

	defer stmt.Close()
	_, err = stmt.Exec(user.ID, item.ID, 1, endDate)
	return
}

func GetCartItems(user User) []CartItem {
	rows, err := db.Query(`
        SELECT Cart.ID, [User ID], [Item ID], Item.Name, Item.Description, Item.Image, Cart.Quantity, [End Date]
        FROM Cart
        INNER JOIN Item ON Cart.[Item ID] = Item.ID
        WHERE [User ID] = $1`, user.ID,
	)
	checkErr(err)

	var items []CartItem

	for rows.Next() {
		var item = CartItem{}
		var endDate string
		err = rows.Scan(&item.ID, &item.UserID, &item.ItemID, &item.Name, &item.Description, &item.Image, &item.Quantity, &endDate)
		checkErr(err)
		parsedDate, err := time.Parse(dbDateLayout, endDate)
		checkErr(err)
		item.EndDate = parsedDate.Format(localDateLayoutInputField)
		items = append(items, item)
	}

	return items
}

func UpdateCartItemQuantity(id int, quantity int) {
	_, err = db.Exec("UPDATE Cart SET Quantity = $1 WHERE ID = $2", quantity, id)
	checkErr(err)
	return
}

func UpdateCartItemDate(id int, date string) {
	parsedDate, err := time.Parse(localDateLayoutInputField, date)
	checkErr(err)
	_, err = db.Exec("UPDATE Cart SET [End Date] = $1 WHERE ID = $2", parsedDate.Format(dbDateLayout), id)
	checkErr(err)
	return
}

func DeleteCartItem(id int) {
	_, err = db.Exec("DELETE FROM Cart WHERE id = $1", id)
	checkErr(err)
	return
}

func DeleteCartItemsByUser(user User) {
	_, err = db.Exec("DELETE FROM Cart WHERE [User ID] = $1", user.ID)
	checkErr(err)
	return
}

func CheckoutCart(user User) {
	items := GetCartItems(user)

	for _, element := range items {
		item, err := GetItem(element.ItemID)
		checkErr(err)
		parsedDate, err := time.Parse(localDateLayoutInputField, element.EndDate)
		checkErr(err)
		CreateEntry(user, item, parsedDate.Format(dbDateLayout), 3, element.Quantity)
	}

	DeleteCartItemsByUser(user)
}

func GetCartItemCount(user User) int {
	var count int
	err = db.QueryRow("SELECT COUNT(ID) FROM Cart WHERE [User ID] = $1", user.ID).Scan(&count)
	checkErr(err)
	return count
}
