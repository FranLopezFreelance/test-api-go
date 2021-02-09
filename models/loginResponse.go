package models

// LoginResponse representa la estructura de la respuesta del login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}