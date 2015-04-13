package blog

import (
	"github.com/etcinit/tumble/entities"
	"github.com/etcinit/tumble/responses"
)

// GetLikesResponse represents a /likes request response
type GetLikesResponse struct {
	Meta     responses.Meta          `json:"meta"`
	Response GetLikesResponseContent `json:"response"`
}

// GetLikesResponseContent represents the inner response content
type GetLikesResponseContent struct {
	LikedPosts []entities.Post `json:"liked_posts"`
	LikedCount int             `json:"liked_count"`
}
