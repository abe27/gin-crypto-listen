package models

type Response struct {
	ID      string      `json:"id"`
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	ID     string `json:"id"`
	Header string `json:"header" default:"Authorization"`
	Type   string `json:"type" default:"Bearer"`
	Token  string `json:"token"`
}
