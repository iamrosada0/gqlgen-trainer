package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

import (
	"context"
	"gqlgen_test/generated"
	"gqlgen_test/model"
)

type Resolver struct{}

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	type Resolver struct{}
*/
