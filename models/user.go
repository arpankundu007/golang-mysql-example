package models

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Scopes []string `json:"scopes"`
}