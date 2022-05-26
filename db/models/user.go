package models

import (
	"github.com/astaxie/beego/orm"

	graphmodel "soccer-manager/graph/model"
)

type User struct {
	Base
	Name     string `json:"name" orm:"column(name)"`
	Email    string `json:"email" orm:"unique;column(email)"`
	Password string `json:"password" orm:"column(password)"`
	Team     *Team  `json:"team" orm:"rel(fk)"`
}

func (*User) TableName() string {
	return "users"
}

func (m *User) Serialize() *graphmodel.User {
	res := &graphmodel.User{
		ID:        m.ID,
		Name:      &m.Name,
		Email:     &m.Email,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
	if m.Team != nil {
		res.Team = m.Team.Serialize()
	}

	return res
}

type UserQuery struct {
	User
}

func init() {
	orm.RegisterModel(new(User))
}
