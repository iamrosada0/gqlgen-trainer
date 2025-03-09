package resolver

import "gqlgen_test/generated"

type Resolver struct {
}

// Mutation retorna o resolver de Mutation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query retorna o resolver de Query.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription retorna o resolver de Subscription.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
