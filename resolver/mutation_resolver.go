package resolver

import (
	"context"
	"gqlgen_test/model"
)

var posts []*model.Post

// MutationResolver para criar um novo post
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string) (*model.Post, error) {
	newPost := &model.Post{
		ID:      "2", // Normalmente, você geraria um UUID
		Title:   title,
		Content: content,
		Author:  &model.User{ID: "1", Name: "Alice"},
	}

	posts = append(posts, newPost)
	return newPost, nil
}
