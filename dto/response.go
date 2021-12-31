package dto

type Response struct {
	Message string `json:"message"`
}

type JWT struct {
	Token string `json:"token"`
}
