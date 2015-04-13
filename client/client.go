package client

import (
	"github.com/etcinit/tumble/client/blog"
	"github.com/etcinit/tumble/client/util"
)

// TumblrClient is an implementation of a Tumblr API client in Go
type TumblrClient struct {
	builder *util.Builder
	Blog    *blog.Subclient
}

// New creates a new instance of a client
func New() *TumblrClient {
	builder := util.NewBuilder("", "")

	client := TumblrClient{
		builder: builder,
	}
	client.setup()

	return &client
}

// NewWithKey creates a new instance of a client with a predefined API key
func NewWithKey(apiKey string) *TumblrClient {
	builder := util.NewBuilder(apiKey, "")

	client := TumblrClient{
		builder: builder,
	}
	client.setup()

	return &client
}

// setup initializes all the internal components of the client
func (t *TumblrClient) setup() {
	t.Blog = blog.New(t.builder)
}
