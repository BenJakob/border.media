package model

// Dummy Data
// Will be replaced with Database Access soon

func GetCarouselImages() (string, []string) {
	return "/img/video_camera.png", []string{"/img/speaker.jpg", "/img/microphone.jpg", "/img/light.jpg", "/img/turntable.jpg", "/img/headset.jpg"}
}

func GetCategories() []string {
	return []string{"Alle", "Kameras", "Beleuchtung", "Monitore", "Sonstiges"}
}

func GetSortOptions() []string {
	return []string{"", "Preis", "Verfügbarkeit", "Kategorie"}
}

func GetEquipment() []Item {
	var items []Item
	for i := 1; i < 9; i++ {
		items = append(items, Item{
			ID:          uint32(i),
			Name:        "Kamera",
			Description: "Eine Kamera ist eine fototechnische Apparatur, die statische oder bewegte Bilder auf einem fotografischen Film oder elektronisch auf ein magnetisches Videoband oder digitales Speichermedium aufzeichnen oder über eine Schnittstelle übermitteln kann.",
			Status:      GetEquipmentStatus(i),
			Image:       "/img/video_camera.png",
		})
	}
	return items
}

func GetEquipmentStatus(id int) string {
	return "entliehen"
}

func GetUsers() []User {
	var users []User
	for i := 1; i < 9; i++ {
		users = append(users, User{
			ID:       uint32(i),
			Name:     "Max",
			LastName: "Mustermann",
			UserName: "Max123",
			Email:    "max123@email.com",
			Password: "12345",
			Image:    "/img/user.svg",
		})
	}
	return users
}

func GetUser() User {
	return User{
		ID:       1,
		Name:     "Max",
		LastName: "Mustermann",
		UserName: "Max123",
		Email:    "max123@email.com",
		Password: "12345",
		Image:    "/img/user.svg",
		Status:   "xxxxxxx",
		Admin:    false,
	}
}

func GetAdmin() User {
	return User{
		ID:       1,
		Name:     "Max",
		LastName: "Mustermann",
		UserName: "Max123",
		Email:    "max123@email.com",
		Password: "12345",
		Image:    "/img/user.svg",
		Status:   "xxxxxxx",
		Admin:    true,
	}
}
