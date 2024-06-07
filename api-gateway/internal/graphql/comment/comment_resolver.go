package comment

import (
	"context"
	"strconv"

	"github.com/frhnfrnk/blog-platform-microservices/comment-service/pb"
	"github.com/graphql-go/graphql"
)

// Resolver struct
type Resolver struct {
	commentServiceClient pb.CommentServiceClient
}

// NewResolver creates a new Resolver
func NewResolver(commentServiceClient pb.CommentServiceClient) *Resolver {
	return &Resolver{
		commentServiceClient: commentServiceClient,
	}
}

func (r *Resolver) resolveComments(p graphql.ResolveParams) (interface{}, error) {
	response, err := r.commentServiceClient.GetAllComments(context.Background(), &pb.GetAllCommentsRequest{})
	if err != nil {
		return nil, err
	}
	return response.Comments, nil
}

func (r *Resolver) resolveCommentByID(params graphql.ResolveParams) (interface{}, error) {
	id, _ := params.Args["id"].(string)
	idInt, err := strconv.Atoi(id)

	comment, err := r.commentServiceClient.GetComment(context.Background(), &pb.GetCommentRequest{Id: uint32(idInt)})
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *Resolver) createComment(p graphql.ResolveParams) (interface{}, error) {
	content, _ := p.Args["content"].(string)
	postID, _ := p.Args["post_id"].(string)
	userID, _ := p.Args["user_id"].(string)

	response, err := r.commentServiceClient.CreateComment(context.Background(), &pb.CreateCommentRequest{
		Content: content,
		PostId:  postID,
		UserId:  userID,
	})
	if err != nil {
		return nil, err
	}
	return response.Comment, nil
}

func (r *Resolver) updateComment(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)
	content, _ := p.Args["content"].(string)

	response, err := r.commentServiceClient.UpdateComment(context.Background(), &pb.UpdateCommentRequest{
		Id:      id,
		Content: content,
	})
	if err != nil {
		return nil, err
	}
	return response.Comment, nil
}

func (r *Resolver) deleteComment(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(string)

	response, err := r.commentServiceClient.DeleteComment(context.Background(), &pb.DeleteCommentRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return response.Message, nil
}
