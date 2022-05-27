package models

import (
	"github.com/astaxie/beego/orm"

	graphmodel "soccer-manager/graph/model"
)

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

func (m *Team) Serialize() *graphmodel.Team {
	res := &graphmodel.Team{
		ID:        m.ID,
		Name:      &m.Name,
		Country:   m.Country,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Budget: &graphmodel.TeamBudget{
			RemainingInDollars: &m.RemainingBudgetInDollars,
		},
	}
	return res
}

type TeamQuery struct {
	Team
}

func init() {
	orm.RegisterModel(new(Team))
}
