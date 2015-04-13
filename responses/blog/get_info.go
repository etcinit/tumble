package blog

import (
	"github.com/etcinit/tumble/entities"
	"github.com/etcinit/tumble/responses"
)

// GetInfoResponse represents a /info request response
type GetInfoResponse struct {
	Meta     responses.Meta         `json:"meta"`
	Response GetInfoResponseContent `json:"response"`
}

// GetInfoResponseContent represents the inner response content
type GetInfoResponseContent struct {
	Blog entities.Blog `json:"blog"`
}
