package util

import (
	"encoding/json"

	"github.com/google/go-querystring/query"
	"github.com/parnurzeal/gorequest"
)

// BuilderInterface should be capable of constructing base requests for the
// client
type BuilderInterface interface {
	Get(path string) *gorequest.SuperAgent
	GetWithParams(path string, params interface{}) (*gorequest.SuperAgent, error)
	GetComplete(path string, params interface{}, response interface{}) []error
	MakeURL(path string) string
	//MakeOAuth() // One day
}

// Builder constructs the base of all requests for the client
type Builder struct {
	consumerKey    string
	consumerSecret string
	superAgent     *gorequest.SuperAgent
}

// NewBuilder constructs a new instance of a builder
func NewBuilder(key string, secret string) *Builder {
	return &Builder{
		consumerKey:    key,
		consumerSecret: secret,
		superAgent:     gorequest.New(),
	}
}

// Get constructs a new request with proper default variables, such as an API
// key parameter for authentication.
func (b *Builder) Get(path string) *gorequest.SuperAgent {
	if b.consumerKey != "" {
		return b.superAgent.Get(path).Query("api_key=" + b.consumerKey)
	}

	return b.superAgent.Get(path)
}

// GetWithParams is just like Get, but it also includes some parameters as part
// of the query string
func (b *Builder) GetWithParams(path string, params interface{}) (*gorequest.SuperAgent, error) {
	query, queryError := query.Values(params)

	if queryError != nil {
		return nil, queryError
	}

	return b.Get(path).Query(query.Encode()), nil
}

// GetComplete performs all the steps needed for a simple GET request with
// optional parameters.
func (b *Builder) GetComplete(path string, params interface{}, response interface{}) []error {
	url := b.MakeURL(path)

	// Decide wether to provide additional parameters or not
	var agent *gorequest.SuperAgent
	if params != nil {
		// Parse struct into a query string and generate the agent
		var paramError error
		agent, paramError = b.GetWithParams(url, params)
		if paramError != nil {
			return []error{paramError}
		}
	} else {
		// Generate a simple agent
		agent = b.Get(url)
	}

	// Perform the requests
	_, body, errors := agent.End()

	// Check if we got any errors
	if len(errors) > 0 {
		return errors
	}

	// Attempt to parse the response
	parseError := json.Unmarshal([]byte(body), response)

	// Check if we got any errors
	if parseError != nil {
		return []error{parseError}
	}

	// Success, no errors
	return nil
}

// MakeURL builds an API request url
func (b *Builder) MakeURL(path string) string {
	return "http://api.tumblr.com/v2/" + path
}
