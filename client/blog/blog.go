package blog

import (
	"encoding/json"

	"github.com/etcinit/tumble/client/util"
	"github.com/etcinit/tumble/responses/blog"
)

// Subclient is a mini-client focused on Blog-related API calls
type Subclient struct {
	builder util.BuilderInterface
}

// New constructs a new instance of the subclient
func New(builder util.BuilderInterface) *Subclient {
	return &Subclient{
		builder: builder,
	}
}

// GetInfo gets information about a single blog
func (b *Subclient) GetInfo(baseHostname string) (blog.GetInfoResponse, []error) {
	url := b.builder.MakeURL("blog/" + baseHostname + "/info")

	_, body, errors := b.builder.Get(url).End()

	parsed := blog.GetInfoResponse{}

	if len(errors) > 0 {
		return parsed, errors
	}

	json.Unmarshal([]byte(body), &parsed)

	return parsed, nil
}

// GetLikes gets a blog's liked posts
func (b *Subclient) GetLikes(baseHostname string, params GetLikesParameters) (blog.GetLikesResponse, []error) {
	parsed := blog.GetLikesResponse{}
	errors := b.builder.GetComplete(
		"blog/"+baseHostname+"/likes",
		params,
		&parsed,
	)

	if len(errors) > 0 {
		return parsed, errors
	}

	return parsed, nil
}

// GetPosts gets a blog's posts
func (b *Subclient) GetPosts(baseHostname string, params GetPostsParameters) (blog.GetPostsResponse, []error) {
	parsed := blog.GetPostsResponse{}
	errors := b.builder.GetComplete(
		"blog/"+baseHostname+"/posts",
		params,
		&parsed,
	)

	if len(errors) > 0 {
		return parsed, errors
	}

	return parsed, nil
}
