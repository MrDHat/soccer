package graphmodel

type Team struct {
	// The id of the team
	ID int64 `json:"id"`
	// The created at timestamp of the team in unix epoch format
	CreatedAt *int64 `json:"createdAt"`
	// The updated at timestamp of the team in unix epoch format
	UpdatedAt *int64 `json:"updatedAt"`
	// The name of the team
	Name *string `json:"name"`
	// The country of the team
	Country *string `json:"country"`
	// The budget details of the team
	Budget *TeamBudget `json:"budget"`
	// The user who owns the team
	User *User `json:"user"`
}
