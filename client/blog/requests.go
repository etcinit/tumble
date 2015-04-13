package blog

// GetLikesParameters are the parameters than can be provided to a GET /likes
// request.
type GetLikesParameters struct {
	Limit  int   `url:"limit,omitempty"`
	Offset int   `url:"offset,omitempty"`
	Before int64 `url:"before,omitempty"`
	After  int64 `url:"after,omitempty"`
}

// GetPostsParameters are the parameters that can be provided to a GET /posts
// request.
type GetPostsParameters struct {
	Type       string `url:"type,omitempty"`
	ID         int64  `url:"id,omitempty"`
	Tag        string `url:"tag,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
	ReblogInfo bool   `url:"reblog_info,omitempty"`
	NotesInfo  bool   `url:"notes_info,omitempty"`
	Filter     string `url:"filter,omitempty"`
}
