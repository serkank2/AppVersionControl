package dto

type RegisterDto struct {
	Error          string           `json:"error"`
	HttpStatusCode int              `json:"httpStatusCode"`
	Result         *RegisterDataDto `json:"result"`
}

type RegisterDataDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Username string `json:"username"`
	Email    string `json:"email" validate:"required,email"`
}
