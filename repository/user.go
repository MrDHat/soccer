package repository

type UserRepo interface {
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}
