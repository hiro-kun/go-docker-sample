package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ono-hiroshi/gqlgen-todos/graph/generated"
	"github.com/ono-hiroshi/gqlgen-todos/graph/model"
)

func (r *mutationResolver) CreateStudent(ctx context.Context, input *model.Student) (*model.Students, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetStudents(ctx context.Context) ([]*model.Students, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
