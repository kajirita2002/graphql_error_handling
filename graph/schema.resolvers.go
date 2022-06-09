package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/http"

	"github.com/kajirita2002/graphql_error_handling/graph/generated"
	"github.com/kajirita2002/graphql_error_handling/graph/model"
)

type AppError struct {
	Code int
	Msg  string
}

var ErrNotFound = AppError{Code: http.StatusNotFound, Msg: "not found"}

func (e AppError) Error() string {
	return fmt.Sprintf("")
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return nil, ErrNotFound
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
