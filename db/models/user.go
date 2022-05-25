package models

type User struct {
	Base
	Name     string `json:"name" orm:"column(name)"`
	Email    string `json:"email" orm:"unique;column(email)"`
	Password string `json:"password" orm:"column(password)"`
	Team     *Team  `json:"team" orm:"rel(fk)"`
}
