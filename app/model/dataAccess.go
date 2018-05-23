package model

// Dummy Data
// Will be replaced with Database Access soon

func GetMinNavbar() Navbar {
	n1 := NavItem{Title: "Equipment", Link: "/equipment"}
	n2 := NavItem{Title: "Login", Link: "/login"}
	return Navbar{Items: []NavItem{n1, n2}, IsLoggedIn: false}
}

func GetNavbarUser() Navbar {
	n1 := NavItem{Title: "Equipment", Link: "/equipment"}
	n2 := NavItem{Title: "Meine Geräte", Link: "/my-equipment"}
	n3 := NavItem{Title: "Logout", Link: "/login"}
	return Navbar{Items: []NavItem{n1, n2, n3}, IsLoggedIn: true}
}

func GetNavbarAdmin() Navbar {
	n1 := NavItem{Title: "Equipment", Link: "/equipment"}
	n2 := NavItem{Title: "Meine Geräte", Link: "/my-equipment"}
	n3 := NavItem{Title: "Logout", Link: "/login"}
	return Navbar{Items: []NavItem{n1, n2, n3}, IsLoggedIn: true}
}

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
