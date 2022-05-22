/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package graphql

import (
	"context"

	"github.com/durudex/durudex-gateway/internal/domain"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GraphQL directive for user authorization.
func (h *Handler) isAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	if ctx.Value(domain.UserCtx) == nil {
		return nil, &gqlerror.Error{Message: "Authorization token failed"}
	}

	return next(ctx)
}
