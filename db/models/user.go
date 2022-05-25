package models

type User struct {
	Base
	Email    string `json:"email" orm:"unique;column(email)"`
	Password string `json:"password" orm:"column(password)"`
	Name     string `json:"name" orm:"column(name)"`
	Team     *Team  `json:"team" orm:"rel(fk)"`
}
