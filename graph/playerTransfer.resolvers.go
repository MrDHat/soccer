package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"soccer-manager/graph/generated"
	graphmodel "soccer-manager/graph/model"
)

func (r *playerTransferResolver) Player(ctx context.Context, obj *graphmodel.PlayerTransfer) (*graphmodel.Player, error) {
	return r.Services.Player().GetForPlayerTransfer(ctx, obj)
}

func (r *playerTransferResolver) OwnerTeam(ctx context.Context, obj *graphmodel.PlayerTransfer) (*graphmodel.Team, error) {
	return r.Services.Team().GetForPlayerTransfer(ctx, obj)
}

// PlayerTransfer returns generated.PlayerTransferResolver implementation.
func (r *Resolver) PlayerTransfer() generated.PlayerTransferResolver {
	return &playerTransferResolver{r}
}

type playerTransferResolver struct{ *Resolver }
