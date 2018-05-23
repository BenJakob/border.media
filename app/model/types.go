package model

type User struct {
	ID       uint32
	UserName string
	Email    string
	Password string
	Image    string
}

type Item struct {
	ID          uint32
	Name        string
	Description string
	Status      string
	Image       string
}

type Header struct {
	Title string
	Css   []string
}

type Navbar struct {
	Items      []NavItem
	IsLoggedIn bool
}

type NavItem struct {
	Title string
	Link  string
}
