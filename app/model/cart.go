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
	statement := "INSERT INTO [Ledger Entry](Type, [User ID], [Item ID], Quantity, [Start Date], [End Date]) VALUES ($1, $2, $3, $4, $5, $6)"
	stmt, err := db.Prepare(statement)
	checkErr(err)

	startDate := time.Now().Format(dbDateLayout)
	endDate := time.Now().AddDate(0, 0, 1).Format(dbDateLayout)

	defer stmt.Close()
	_, err = stmt.Exec(1, user.ID, item.ID, 1, startDate, endDate)
	return
}

func GetCartItems(user User) []CartItem {
	rows, err := db.Query(`
        SELECT [Ledger Entry].ID, [User ID], [Item ID], Item.Name, Item.Description, Item.Image, [Ledger Entry].Quantity, [End Date]
        FROM [Ledger Entry]
        INNER JOIN Item ON [Ledger Entry].[Item ID] = Item.ID
        WHERE [User ID] = $1 AND [Type] = 1`, user.ID,
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

// func UpdateCartItemQuantity(id int, quantity int) {
// 	_, err = db.Exec("UPDATE Cart SET Quantity = $1 WHERE ID = $2", quantity, id)
// 	checkErr(err)
// 	return
// }

func UpdateCartItemQuantity(id int, quantity int) {
	_, err = db.Exec("UPDATE [Ledger Entry] SET Quantity = $1 WHERE ID = $2", quantity, id)
	checkErr(err)
	return
}

func UpdateCartItemDate(id int, date string) {
	parsedDate, err := time.Parse(localDateLayoutInputField, date)
	checkErr(err)
	_, err = db.Exec("UPDATE [Ledger Entry] SET [End Date] = $1 WHERE ID = $2", parsedDate.Format(dbDateLayout), id)
	checkErr(err)
	return
}

func DeleteCartItem(id int) {
	_, err = db.Exec("DELETE FROM [Ledger Entry] WHERE id = $1", id)
	checkErr(err)
	return
}

func DeleteCartItemsByUser(user User) {
	_, err = db.Exec("DELETE FROM [Ledger Entry] WHERE [User ID] = $1 AND Type = 1", user.ID)
	checkErr(err)
	return
}

func CheckoutCart(user User) {
	_, err = db.Exec("UPDATE [Ledger Entry] SET Type = $1 WHERE [User ID] = $2 AND Type = 1", 3, user.ID)
	checkErr(err)
	return
}

func GetCartItemCount(user User) int {
	var count int
	err = db.QueryRow("SELECT COUNT(ID) FROM [Ledger Entry] WHERE [User ID] = $1 AND Type = 1", user.ID).Scan(&count)
	checkErr(err)
	return count
}
