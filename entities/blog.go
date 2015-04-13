package entities

// Blog represents a blog on Tumblr
type Blog struct {
	Title       string    `json:"title"`
	Posts       int       `json:"posts"`
	Name        string    `json:"name"`
	Updated     int64     `json:"updated"`
	Description string    `json:"description"`
	Ask         bool      `json:"ask"`
	AskAnon     bool      `json:"ask_anon"`
	Likes       int       `json:"likes"`
	Theme       BlogTheme `json:"theme"`
}

// BlogTheme describes the theme of a blog
type BlogTheme struct {
	AvatarShape        string      `json:"avatar_shape"`
	BackgroundColor    string      `json:"background_color"`
	BodyFont           string      `json:"body_font"`
	HeaderBounds       interface{} `json:"header_bounds"`
	HeaderImage        string      `json:"header_image"`
	HeaderImageFocused string      `json:"header_image_focused"`
	HeaderImageScaled  string      `json:"header_image_scaled"`
	HeaderStretch      bool        `json:"header_stretch"`
	LinkColor          string      `json:"link_color"`
	ShowAvatar         bool        `json:"show_avatar"`
	ShowDescription    bool        `json:"show_description"`
	ShowHeaderImage    bool        `json:"show_header_image"`
	ShowTitle          bool        `json:"show_title"`
	TitleColor         string      `json:"title_color"`
	TitleFont          string      `json:"title_font"`
	TitleFontWeight    string      `json:"title_font_weight"`
}
