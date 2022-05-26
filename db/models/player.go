package models

import "github.com/astaxie/beego/orm"

type Player struct {
	Base
	FirstName             string `json:"first_name" orm:"column(first_name)"`
	LastName              string `json:"last_name" orm:"column(last_name)"`
	Age                   int64  `json:"age" orm:"column(age)"`
	CurrentValueInDollars int64  `json:"current_value_in_dollars" orm:"column(current_value_in_dollars)"`
	PlayerType            string `json:"player_type" orm:"column(player_type)"`
	Team                  *Team  `json:"team" orm:"rel(fk)"`
}

func (*Player) TableName() string {
	return "players"
}

func init() {
	orm.RegisterModel(new(Player))
}
