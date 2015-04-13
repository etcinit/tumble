package entities

// Photo represents a photo on a Tumblr blog
type Photo struct {
	Caption      string      `json:"caption"`
	AltSizes     []PhotoSize `json:"alt_sizes"`
	OriginalSize PhotoSize   `json:"original_size"`
	Exif         ExifData    `json:"exif"`
}

// PhotoSize contains dimensions of a photo URL
type PhotoSize struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// ExifData contains information about the camera used to take a photo
type ExifData struct {
	Camera      string `json:"Camera"`
	ISO         int    `json:"ISO"`
	Aperture    string `json:"Aperture"`
	Exposure    string `json:"Exposure"`
	FocalLength string `json:"FocalLength"`
}
