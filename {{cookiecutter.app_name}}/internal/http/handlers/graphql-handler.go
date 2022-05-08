package handlers

import (
	"context"
	"errors"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/gin-gonic/gin"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/ent/privacy"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/graph/resolver"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/config"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Defining the Graphql handler
func GraphqlHandler(client *ent.Client, conf *config.Config, redisCache *graphql.Cache) gin.HandlerFunc {

	srv := handler.NewDefaultServer(resolver.NewSchema(client))

	// set error handler
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		// add more errors here
		if errors.Is(err, privacy.Deny) {
			err.Message = "permission denied"
		}

		return err
	})

	// remove this line if you do not use transactions
	srv.Use(entgql.Transactioner{TxOpener: client})
	//srv.AddTransport(transport.Options{})
	//srv.AddTransport(transport.GET{})
	//srv.AddTransport(transport.POST{})
	//srv.AddTransport(transport.MultipartForm{})
	// srv.Use(extension.FixedComplexityLimit(4))
	if (*conf).Database.Debug {
		srv.Use(&debug.Tracer{})
	}

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
