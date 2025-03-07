package resolver

// THIS CODE WILL BE UPDATED WITH SCHEMA CHANGES. PREVIOUS IMPLEMENTATION FOR SCHEMA CHANGES WILL BE KEPT IN THE COMMENT SECTION. IMPLEMENTATION FOR UNCHANGED SCHEMA WILL BE KEPT.

type Resolver struct{}

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	var posts = []models.Post{
	{
		ID:      "1",
		Title:   "Meu primeiro post",
		Content: "Este é o conteúdo do primeiro post",
		Author: models.User{
			ID:   "1",
			Name: "João",
		},
	},
	{
		ID:      "2",
		Title:   "Segundo post",
		Content: "Este é o conteúdo do segundo post",
		Author: models.User{
			ID:   "2",
			Name: "Maria",
		},
	},
}
type Resolver struct{}
func (r *Resolver) Posts() []*models.Post {
	var result []*models.Post
	for i := range posts {
		result = append(result, &posts[i])
	}
	return result
}
*/
