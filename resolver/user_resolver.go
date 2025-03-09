package resolver

import (
	"context"
	"gqlgen_test/model"
)

// Resolver para a query de usuários
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
	}, nil
}
