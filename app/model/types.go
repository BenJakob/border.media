package model

type User struct {
	ID       uint32
	Name     string
	LastName string
	UserName string
	Email    string
	Password string
	Image    string
	Status   string
	Admin    bool
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
