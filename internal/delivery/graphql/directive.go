package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQL directive for checking email code.
func (h *Handler) emailCode(ctx context.Context, obj interface{}, next graphql.Resolver, email string, code uint64) (interface{}, error) {
	status, err := h.service.Code.CheckByEmail(ctx, email, code)
	if err != nil {
		return nil, err
	}

	// Check codes.
	if !status {
		return nil, &gqlerror.Error{Message: "Wrong code"}
	}

	return next(ctx)
}
