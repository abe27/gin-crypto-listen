package models

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AuthResponse struct {
	Header string `default:"Authorization"`
	Type   string `default:"Bearer"`
	Token  string
}
