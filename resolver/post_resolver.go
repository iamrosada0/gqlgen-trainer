package resolver

import (
	"context"
	"gqlgen_test/model"
)

// Resolver para a query de posts
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return []*model.Post{
		{
			ID:      "1",
			Title:   "GraphQL Introduction",
			Content: "Learning GraphQL...",
		},
	}, nil
}

// Resolver para criar um post
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string) (*model.Post, error) {
	newPost := &model.Post{
		ID:      "2",
		Title:   title,
		Content: content,
	}
	return newPost, nil
}
