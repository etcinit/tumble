package entities

// Post represents a Tumblr blog post
type Post struct {
	BlogName        string      `json:"blog_name"`
	ID              interface{} `json:"id"`
	URL             string      `json:"post_url"`
	Slug            string      `json:"slug"`
	Type            string      `json:"type"`
	Date            string      `json:"date"`
	Timestamp       int64       `json:"timestamp"`
	State           string      `json:"state"`
	Format          string      `json:"format"`
	ReblogKey       string      `json:"reblog_key"`
	Tags            []string    `json:"tags"`
	ShortURL        string      `json:"short_url"`
	Highlighted     []string    `json:"highlighted"`
	NoteCount       int         `json:"note_count"`
	SourceURL       string      `json:"source_url"`
	SourceTitle     string      `json:"source_title"`
	Caption         string      `json:"caption"`
	CaptionAbstract string      `json:"caption_abstract"`
	Body            string      `json:"body"`
	Reblog          Reblog      `json:"reblog"`
	ImagePermalink  string      `json:"image_permalink"`
	PhotosetLayout  string      `json:"photoset_layout"`
	Photos          []Photo     `json:"photos"`
	LikedTimestamp  int64       `json:"liked_timestamp"`
	IsSubmission    bool        `json:"is_submission"`
	PostAuthor      string      `json:"post_author"`
	FeaturedInTag   []string    `json:"featured_in_tag"`
	Artist          string      `json:"artist"`
	Year            int         `json:"year"`
	Album           string      `json:"album"`
	Track           interface{} `json:"track"`
	TrackName       string      `json:"track_name"`
	AlbumArt        string      `json:"album_art"`
	Player          interface{} `json:"player"`
	Embed           string      `json:"embed"`
	Plays           int         `json:"plays"`
	AudioURL        string      `json:"audio_url"`
	AudioSourceURL  string      `json:"audio_source_url"`
	VideoURL        string      `json:"video_url"`
	IsExternal      bool        `json:"is_external"`
	AudioType       string      `json:"audio_type"`
	VideoType       string      `json:"video_type"`
	Duration        int         `json:"duration"`
	HTML5Capable    bool        `json:"html5_capable"`
	ThumbnailURL    string      `json:"thumbnail_url"`
	ThumbnailWidth  int         `json:"thumbnail_width"`
	ThumbnailHeight int         `json:"thumbnail_height"`
}

// Reblog contains information about the rebloged post
type Reblog struct {
	TreeHTML string            `json:"tree_html"`
	Comment  string            `json:"comment"`
	Trail    []ReblogTrailItem `json:"trail"`
}

// ReblogTrailItem is one item in the reblog trail
type ReblogTrailItem struct {
	Blog    Blog   `json:"blog"`
	Post    Post   `json:"post"`
	Comment string `json:"comment"`
}
