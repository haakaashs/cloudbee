package models

type Post struct {
	ID              int32    `json:"id,omitempty"`
	Title           string   `json:"title,omitempty"`
	Content         string   `json:"content,omitempty"`
	Author          string   `json:"author,omitempty"`
	PublicationData string   `json:"-"`
	Tags            []string `json:"tags,omitempty"`
}
