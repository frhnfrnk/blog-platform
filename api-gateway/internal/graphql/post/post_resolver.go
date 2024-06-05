package post

import (
	"context"
	"github.com/frhnfrnk/blog-platform-microservices/post-service/pb"
	"github.com/graphql-go/graphql"
)

// Resolver struct
type Resolver struct {
	postServiceClient pb.PostServiceClient
}

// NewResolver creates a new Resolver
func NewResolver(postServiceClient pb.PostServiceClient) *Resolver {
	return &Resolver{
		postServiceClient: postServiceClient,
	}
}

func (r *Resolver) resolvePosts(p graphql.ResolveParams) (interface{}, error) {
	response, err := r.postServiceClient.GetAllPosts(context.Background(), &pb.GetAllPostsRequest{})
	if err != nil {
		return nil, err
	}
	return response.Posts, nil
}

func (r *Resolver) resolvePostByID(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)

	post, err := r.postServiceClient.GetPost(context.Background(), &pb.GetPostRequest{Id: id})
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *Resolver) createPost(p graphql.ResolveParams) (interface{}, error) {
	title, _ := p.Args["title"].(string)
	content, _ := p.Args["content"].(string)
	authorID, _ := p.Args["authorID"].(string)

	response, err := r.postServiceClient.CreatePost(context.Background(), &pb.CreatePostRequest{Title: title, Content: content, AuthorId: authorID})
	if err != nil {
		return nil, err
	}
	return response.Post, nil
}

func (r *Resolver) deletePost(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)

	response, err := r.postServiceClient.DeletePost(context.Background(), &pb.DeletePostRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return response.Message, nil
}

func (r *Resolver) updatePost(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	title, _ := p.Args["title"].(string)
	content, _ := p.Args["content"].(string)

	response, err := r.postServiceClient.UpdatePost(context.Background(), &pb.UpdatePostRequest{Id: id, Title: title, Content: content})
	if err != nil {
		return nil, err
	}
	return response.Post, nil
}
