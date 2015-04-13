package blog

import (
	"github.com/etcinit/tumble/entities"
	"github.com/etcinit/tumble/responses"
)

// GetPostsResponse represents a /posts request response
type GetPostsResponse struct {
	Meta     responses.Meta          `json:"meta"`
	Response GetPostsResponseContent `json:"response"`
}

// GetPostsResponseContent represents the inner response content
type GetPostsResponseContent struct {
	Blog       entities.Blog   `json:"blog"`
	Posts      []entities.Post `json:"posts"`
	TotalPosts int             `json:"total_posts"`
}
