package models

import "github.com/astaxie/beego/orm"

type Team struct {
	Base
	Name                     string    `json:"name" orm:"unique;column(name)"`
	Country                  *string   `json:"country" orm:"column(country)"`
	RemainingBudgetInDollars int64     `json:"remaining_budget_in_dollars" orm:"column(remaining_budget_in_dollars)"`
	Players                  []*Player `json:"players" orm:"reverse(many)"`
}

func (*Team) TableName() string {
	return "teams"
}

func init() {
	orm.RegisterModel(new(Team))
}
