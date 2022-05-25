package models

type Team struct {
	Base
	Name                     string  `json:"name" orm:"unique;column(name)"`
	Country                  *string `json:"country" orm:"column(country)"`
	RemainingBudgetInDollars int64   `json:"remaining_budget_in_dollars" orm:"column(remaining_budget_in_dollars)"`
}
