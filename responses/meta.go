package responses

// Meta is some infomration about the responses provided on each response from
// the API.
type Meta struct {
	Status  int    `json:"status"`
	Message string `json:"msg"`
}
