package user

type User struct {
	Data []UserData `json:"data"`
}

type UserData struct {
	ID        uint   `json:"ID"`
	CreatedAt string `json:"CreatedAt"`
	DeletedAt string `json:"DeletedAt"`
	UpdatedAt string `json:"UpdatedAt"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	Name      string `json:"Name"`
}

type Large struct {
	Data []LargeData `json:"data"`
}

type LargeData struct {
	A int `json:"a"`
}
