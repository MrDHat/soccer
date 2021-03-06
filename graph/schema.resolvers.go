package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"soccer-manager/graph/generated"
	graphmodel "soccer-manager/graph/model"
)

func (r *mutationResolver) Signup(ctx context.Context, input graphmodel.SignupInput) (*graphmodel.LoginResponse, error) {
	return r.Services.User().Signup(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input graphmodel.LoginInput) (*graphmodel.LoginResponse, error) {
	return r.Services.User().Login(ctx, input)
}

func (r *mutationResolver) UpdateTeam(ctx context.Context, input graphmodel.UpdateTeamInput) (*graphmodel.Team, error) {
	return r.Services.Team().Update(ctx, input)
}

func (r *mutationResolver) UpdatePlayer(ctx context.Context, input graphmodel.UpdatePlayerInput) (*graphmodel.Player, error) {
	return r.Services.Player().Update(ctx, input)
}

func (r *mutationResolver) BuyPlayer(ctx context.Context, input graphmodel.BuyPlayerInput) (bool, error) {
	return r.Services.Transfer().BuyPlayer(ctx, input)
}

func (r *mutationResolver) CreateTransfer(ctx context.Context, input graphmodel.CreateTransferInput) (*graphmodel.PlayerTransfer, error) {
	return r.Services.Transfer().Create(ctx, input)
}

func (r *queryResolver) Me(ctx context.Context) (*graphmodel.User, error) {
	return r.Services.User().Me(ctx)
}

func (r *queryResolver) MyTeam(ctx context.Context) (*graphmodel.Team, error) {
	return r.Services.Team().My(ctx)
}

func (r *queryResolver) PlayerTransfers(ctx context.Context, input *graphmodel.PlayerTransferListInput) (*graphmodel.PlayerTransferList, error) {
	return r.Services.Transfer().List(ctx, input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
