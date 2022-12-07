package domain

type ApiRegister struct {
	User
	ReqPassword       string `json:"password" validate:"required,min=8,max=100"`
	ReqRepeatPassword string `json:"repeatPassword" validate:"required,min=8,max=100,eqfield=ReqPassword"`
}

type ApiLogin struct {
	Username string `json:"username" validate:"required,min=3,max=100"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}
