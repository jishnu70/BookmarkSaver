package models

type CreateBookmarkInput struct {
	Title string   `json:"title" binding:"required"`
	URL   string   `json:"url" binding:"required"`
	Tags  []string `json:"tags"`
}
