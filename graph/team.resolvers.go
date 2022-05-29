package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"soccer-manager/graph/generated"
	graphmodel "soccer-manager/graph/model"
)

func (r *teamResolver) ValueInDollars(ctx context.Context, obj *graphmodel.Team) (*int64, error) {
	return r.Services.Team().ValueInDollars(ctx, obj)
}

func (r *teamResolver) Players(ctx context.Context, obj *graphmodel.Team, input *graphmodel.TeamPlayerListInput) (*graphmodel.PlayerList, error) {
	return r.Services.Player().ListForTeam(ctx, obj, input)
}

// Team returns generated.TeamResolver implementation.
func (r *Resolver) Team() generated.TeamResolver { return &teamResolver{r} }

type teamResolver struct{ *Resolver }
