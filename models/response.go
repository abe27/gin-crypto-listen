package models

type Response struct {
	ID      string      `form:"id"`
	Success bool        `form:"success" default:"true"`
	Message string      `form:"message"`
	Data    interface{} `form:"data"`
}

type AuthResponse struct {
	ID     string `form:"id"`
	Header string `form:"header" default:"Authorization"`
	Type   string `form:"type" default:"Bearer"`
	Token  string `form:"token"`
}
