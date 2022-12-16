package models

type LoginModel struct {
	EmailOrUsername string `json:"emailOrUsername"`
	Password        string `json:"password"`
}
