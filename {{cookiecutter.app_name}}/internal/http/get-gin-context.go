package http

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve server context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("server context has wrong type")
		return nil, err
	}
	return gc, nil
}
