package models

type UpdateBookMarkInput struct {
	Title *string  `json:"title"`
	URL   *string  `json:"url"`
	Tags  []string `json:"tags"`
}
