package model

type AdminData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	FooterData []string
}

type AdminAddData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	Categories []string
	Locations  []string
	FooterData []string
}

type AdminEditEquipmentData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	Item       Item
	Categories []string
	Locations  []string
	FooterData []string
}

type AdminClientsData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	Users      []User
	UserStatus []string
	FooterData []string
}

type AdminEditClientData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	Client     User
	FooterData []string
}

type AdminEquipmentData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	Categories []string
	Items      []Item
	FooterData []string
}

type CartData struct {
	HeaderData    Header
	User          User
	IsLoggedIn    bool
	Items         []CartItem
	CartItemCount int
	FooterData    []string
}

type EquipmentData struct {
	HeaderData    Header
	User          User
	IsLoggedIn    bool
	Categories    []string
	SortOptions   []string
	Rows          []Row
	CartItemCount int
	FooterData    []string
}

type IndexData struct {
	HeaderData    Header
	User          User
	IsLoggedIn    bool
	ActiveImg     string
	CartItemCount int
	Images        []string
	FooterData    []string
}

type Header struct {
	Title string
	Css   []string
}

type LoginData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	LoginError bool
	FooterData []string
}

type MyEquipmentData struct {
	HeaderData    Header
	User          User
	IsLoggedIn    bool
	RentedItems   []LedgerEntry
	MarkedItems   []LedgerEntry
	CartItemCount int
	FooterData    []string
}

type ProfilData struct {
	HeaderData    Header
	User          User
	IsLoggedIn    bool
	LoginError    bool
	CartItemCount int
	FooterData    []string
}

type RegisterData struct {
	HeaderData Header
	User       User
	IsLoggedIn bool
	LoginError bool
	FooterData []string
}

type Row struct {
	Items []Item
}
