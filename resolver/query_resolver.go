package resolver

import (
	"context"
	"gqlgen_test/model"
)

// QueryResolver retorna a lista de usuários
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{ID: "1", Name: "Alice", Post: nil},
		{ID: "2", Name: "Bob", Post: nil},
	}, nil
}

// QueryResolver retorna a lista de posts
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return []*model.Post{
		{
			ID:      "1",
			Title:   "Meu primeiro post",
			Content: "Este é um conteúdo de teste",
			Author:  &model.User{ID: "1", Name: "Alice"},
		},
	}, nil
}
