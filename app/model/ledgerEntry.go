package model

import "time"

type LedgerEntry struct {
	ID              string
	UserID          uint32
	UserName        string
	ItemID          uint32
	ItemName        string
	ItemDescription string
	ItemImage       string
	StartDate       string
	EndDate         string
	Quantity        int
}

func CreateEntry(user User, item Item, endDate string, entryType int, quantity int) {
	statement := "INSERT INTO [Ledger Entry](Type, [User ID], [Item ID], [Start Date], [End Date], Quantity) values ($1, $2, $3, $4, $5, $6)"
	stmt, err := db.Prepare(statement)

	if err != nil {
		return
	}

	startDate := time.Now().Format(dbDateLayout)

	defer stmt.Close()
	_, err = stmt.Exec(entryType, user.ID, item.ID, startDate, endDate, quantity)
	return
}

func EntryExists(id int) bool {
	entry := LedgerEntry{}
	err := db.QueryRow(`
		SELECT [Ledger Entry].ID
		FROM [Ledger Entry]
		WHERE ID = $1`, id).Scan(&entry.ID)

	return err == nil
}

func getEntry(id int) (LedgerEntry, error) {
	entry := LedgerEntry{}
	err := db.QueryRow(`
		SELECT [Ledger Entry].ID, [Item ID], [Start Date], [End Date], Quantity
		FROM [Ledger Entry]
		WHERE ID = $1`, id).Scan(&entry.ID, &entry.ItemID, &entry.StartDate, &entry.EndDate, &entry.Quantity)
	return entry, err
}

func GetEntriesByCustomer(user User) ([]LedgerEntry, []LedgerEntry) {
	rentedItems := getEntriesByUser(user, 3)
	markedItems := getEntriesByUser(user, 2)
	return rentedItems, markedItems
}

func getEntriesByUser(user User, entryType int) []LedgerEntry {
	timeNow := time.Now().Format(dbDateLayout)
	rows, err := db.Query(`
		SELECT [Ledger Entry].ID, Item.ID, Item.Name, Item.Description, Item.Image, [Start Date], [End Date], [Ledger Entry].Quantity
		FROM [Ledger Entry]
		INNER JOIN Item on [Item ID] = Item.ID AND Type = $1 AND [End Date] >= $2
		WHERE [User ID] = $3`, entryType, timeNow, user.ID,
	)
	checkErr(err)

	var items []LedgerEntry

	for rows.Next() {
		var entry = LedgerEntry{}
		err = rows.Scan(&entry.ID, &entry.ItemID, &entry.ItemName, &entry.ItemDescription, &entry.ItemImage, &entry.StartDate, &entry.EndDate, &entry.Quantity)
		checkErr(err)

		parsedDate, err := time.Parse(dbDateLayout, entry.StartDate)
		checkErr(err)
		entry.StartDate = parsedDate.Format(localDateLayout)

		parsedDate, err = time.Parse(dbDateLayout, entry.EndDate)
		checkErr(err)
		entry.EndDate = parsedDate.Format(localDateLayout)

		items = append(items, entry)
	}
	return items
}

func getAllEntriesByUser(user User, entryType int) []LedgerEntry {
	rows, err := db.Query(`
		SELECT [Ledger Entry].ID, Item.ID, Item.Name, Item.Description, Item.Image, [Start Date], [End Date], [Ledger Entry].Quantity
		FROM [Ledger Entry]
		INNER JOIN Item on [Item ID] = Item.ID AND Type = $1
		WHERE [User ID] = $3`, entryType, user.ID,
	)
	checkErr(err)

	var items []LedgerEntry

	for rows.Next() {
		var entry = LedgerEntry{}
		err = rows.Scan(&entry.ID, &entry.ItemID, &entry.ItemName, &entry.ItemDescription, &entry.ItemImage, &entry.StartDate, &entry.EndDate, &entry.Quantity)
		checkErr(err)

		parsedDate, err := time.Parse(dbDateLayout, entry.StartDate)
		checkErr(err)
		entry.StartDate = parsedDate.Format(localDateLayout)

		parsedDate, err = time.Parse(dbDateLayout, entry.EndDate)
		checkErr(err)
		entry.EndDate = parsedDate.Format(localDateLayout)

		items = append(items, entry)
	}
	return items
}

func getEntriesByItem(item Item, entryType int) []LedgerEntry {
	timeNow := time.Now().Format(dbDateLayout)
	rows, err := db.Query(`
		SELECT [Ledger Entry].ID, User.ID, User.[User Name], Item.ID, Item.Name, Item.Description, Item.Image, [Start Date], [End Date], [Ledger Entry].Quantity
		FROM [Ledger Entry]
		INNER JOIN Item on [Item ID] = Item.ID AND Type = $1 AND [End Date] >= $2
		INNER JOIN User on [User ID] = User.ID
		WHERE [Item ID] = $3`, entryType, timeNow, item.ID,
	)
	checkErr(err)

	var items []LedgerEntry

	for rows.Next() {
		var entry = LedgerEntry{}
		err = rows.Scan(&entry.ID, &entry.UserID, &entry.UserName, &entry.ItemID, &entry.ItemName, &entry.ItemDescription, &entry.ItemImage, &entry.StartDate, &entry.EndDate, &entry.Quantity)
		checkErr(err)

		parsedDate, err := time.Parse(dbDateLayout, entry.StartDate)
		checkErr(err)
		entry.StartDate = parsedDate.Format(localDateLayout)

		parsedDate, err = time.Parse(dbDateLayout, entry.EndDate)
		checkErr(err)
		entry.EndDate = parsedDate.Format(localDateLayout)

		items = append(items, entry)
	}
	return items
}

func ExtendEntry(id int) {
	entry, err := getEntry(id)
	checkErr(err)
	parsedDate, err := time.Parse(dbDateLayout, entry.EndDate)
	checkErr(err)
	newEndDate := parsedDate.AddDate(0, 0, 7).Format(dbDateLayout)
	_, err = db.Exec("UPDATE [Ledger Entry] SET [End Date] = $1 where id = $2", newEndDate, id)
	checkErr(err)
	return
}

func DeleteEntry(id int) {
	_, err = db.Exec("DELETE from [Ledger Entry] where id = $1", id)
	checkErr(err)
	return
}
