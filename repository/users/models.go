package users

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Permissions []Permission
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
