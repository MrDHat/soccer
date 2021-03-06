// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphmodel

import (
	"fmt"
	"io"
	"strconv"
)

type BuyPlayerInput struct {
	PlayerTransferID int64 `json:"playerTransferId"`
}

type CreateTransferInput struct {
	// The id of the player
	PlayerID int64 `json:"playerId"`
	// The amount in dollars that the player will be transferred for
	AmountInDollars int64 `json:"amountInDollars"`
}

type LoginInput struct {
	// The email id of the user which is used for login
	Email string `json:"email"`
	// The password of the user
	Password string `json:"password"`
}

type LoginResponse struct {
	// The token of the user
	Token string `json:"token"`
	// The user
	User *User `json:"user"`
}

type PaginationInput struct {
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type Player struct {
	// The id of the player
	ID int64 `json:"id"`
	// The created at timestamp of the player in unix epoch format
	CreatedAt *int64 `json:"createdAt"`
	// The updated at timestamp of the player in unix epoch format
	UpdatedAt *int64 `json:"updatedAt"`
	// The first name of the player
	FirstName *string `json:"firstName"`
	// The last name of the player
	LastName *string `json:"lastName"`
	// The age of the player
	Age *int64 `json:"age"`
	// Current value of the player in dollars
	CurrentValueInDollars *int64 `json:"currentValueInDollars"`
	// Type of the player
	Type *PlayerType `json:"type"`
	// The country of the player
	Country *string `json:"country"`
	// The transfer status of the player
	TransferStatus *TransferStatus `json:"transferStatus"`
	// The team that the player belongs to
	Team *Team `json:"team"`
}

type PlayerList struct {
	TotalPage    *int64    `json:"totalPage"`
	CurrentPage  *int64    `json:"currentPage"`
	TotalRecords *int64    `json:"totalRecords"`
	Data         []*Player `json:"data"`
}

type PlayerTransferList struct {
	TotalPage    *int64            `json:"totalPage"`
	CurrentPage  *int64            `json:"currentPage"`
	TotalRecords *int64            `json:"totalRecords"`
	Data         []*PlayerTransfer `json:"data"`
}

type PlayerTransferListInput struct {
	OnlyMine   *bool                 `json:"onlyMine"`
	Status     *PlayerTransferStatus `json:"status"`
	Pagination *PaginationInput      `json:"pagination"`
}

type SignupInput struct {
	// The email id of the user which is used for login
	Email string `json:"email"`
	// The name of the user
	Name string `json:"name"`
	// The password of the user
	Password string `json:"password"`
}

type TeamBudget struct {
	// The amount of money the team has currently
	RemainingInDollars *int64 `json:"remainingInDollars"`
}

type TeamPlayerListInput struct {
	Pagination *PaginationInput `json:"pagination"`
}

type UpdatePlayerInput struct {
	// The id of the player
	ID int64 `json:"id"`
	// The first name of the player
	FirstName *string `json:"firstName"`
	// The last name of the player
	LastName *string `json:"lastName"`
	// The country of the player
	Country *string `json:"country"`
}

type UpdateTeamInput struct {
	// The id of the team
	ID int64 `json:"id"`
	// The name of the team
	Name *string `json:"name"`
	// The country of the team
	Country *string `json:"country"`
}

type UpdateUserInput struct {
	// The id of the user
	ID int64 `json:"id"`
	// The name of the user
	Name string `json:"name"`
}

type User struct {
	// The id of the user
	ID int64 `json:"id"`
	// The created at timestamp of the user in unix epoch format
	CreatedAt *int64 `json:"createdAt"`
	// The updated at timestamp of the user in unix epoch format
	UpdatedAt *int64 `json:"updatedAt"`
	// The email id of the user which is used for login
	Email *string `json:"email"`
	// The name of the user
	Name *string `json:"name"`
	// The team this user belongs to
	Team *Team `json:"team"`
}

type PlayerTransferStatus string

const (
	PlayerTransferStatusPending   PlayerTransferStatus = "pending"
	PlayerTransferStatusCompleted PlayerTransferStatus = "completed"
)

var AllPlayerTransferStatus = []PlayerTransferStatus{
	PlayerTransferStatusPending,
	PlayerTransferStatusCompleted,
}

func (e PlayerTransferStatus) IsValid() bool {
	switch e {
	case PlayerTransferStatusPending, PlayerTransferStatusCompleted:
		return true
	}
	return false
}

func (e PlayerTransferStatus) String() string {
	return string(e)
}

func (e *PlayerTransferStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlayerTransferStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlayerTransferStatus", str)
	}
	return nil
}

func (e PlayerTransferStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PlayerType string

const (
	PlayerTypeGoalkeeper PlayerType = "goalkeeper"
	PlayerTypeDefender   PlayerType = "defender"
	PlayerTypeMidfielder PlayerType = "midfielder"
	PlayerTypeAttacker   PlayerType = "attacker"
)

var AllPlayerType = []PlayerType{
	PlayerTypeGoalkeeper,
	PlayerTypeDefender,
	PlayerTypeMidfielder,
	PlayerTypeAttacker,
}

func (e PlayerType) IsValid() bool {
	switch e {
	case PlayerTypeGoalkeeper, PlayerTypeDefender, PlayerTypeMidfielder, PlayerTypeAttacker:
		return true
	}
	return false
}

func (e PlayerType) String() string {
	return string(e)
}

func (e *PlayerType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PlayerType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PlayerType", str)
	}
	return nil
}

func (e PlayerType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortBy string

const (
	SortByLatest SortBy = "latest"
	SortByOldest SortBy = "oldest"
)

var AllSortBy = []SortBy{
	SortByLatest,
	SortByOldest,
}

func (e SortBy) IsValid() bool {
	switch e {
	case SortByLatest, SortByOldest:
		return true
	}
	return false
}

func (e SortBy) String() string {
	return string(e)
}

func (e *SortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortBy", str)
	}
	return nil
}

func (e SortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SortOrder string

const (
	SortOrderAsc  SortOrder = "asc"
	SortOrderDesc SortOrder = "desc"
)

var AllSortOrder = []SortOrder{
	SortOrderAsc,
	SortOrderDesc,
}

func (e SortOrder) IsValid() bool {
	switch e {
	case SortOrderAsc, SortOrderDesc:
		return true
	}
	return false
}

func (e SortOrder) String() string {
	return string(e)
}

func (e *SortOrder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortOrder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortOrder", str)
	}
	return nil
}

func (e SortOrder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransferStatus string

const (
	TransferStatusOwned  TransferStatus = "owned"
	TransferStatusOnSale TransferStatus = "on_sale"
)

var AllTransferStatus = []TransferStatus{
	TransferStatusOwned,
	TransferStatusOnSale,
}

func (e TransferStatus) IsValid() bool {
	switch e {
	case TransferStatusOwned, TransferStatusOnSale:
		return true
	}
	return false
}

func (e TransferStatus) String() string {
	return string(e)
}

func (e *TransferStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransferStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransferStatus", str)
	}
	return nil
}

func (e TransferStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
