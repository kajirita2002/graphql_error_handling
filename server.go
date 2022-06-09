package main

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kajirita2002/graphql_error_handling/graph"
	"github.com/kajirita2002/graphql_error_handling/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var appErr graph.AppError
		if !errors.As(err, &appErr) {
			return err
		}
		return &gqlerror.Error{
			Message: appErr.Msg,
			Path:    graphql.GetPath(ctx),
			Extensions: map[string]interface{}{
				"code": appErr.Code,
			},
		}
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		return &gqlerror.Error{
			Extensions: map[string]interface{}{
				"code":  http.StatusInternalServerError,
				"cause": err,
			},
		}
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
