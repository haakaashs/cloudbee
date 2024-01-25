package models

type Post struct {
	ID              int32    `json:"id,omitempty"`
	Title           string   `json:"title,omitempty"`
	Content         string   `json:"content,omitempty"`
	Author          string   `json:"author,omitempty"`
	PublicationData string   `json:"-"`
	Tags            []string `json:"tags,omitempty"`
}

// PostId          int32    `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
// Title           string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
// Content         string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
// Author          string   `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
// PublicationData string   `protobuf:"bytes,5,opt,name=publication_data,json=publicationData,proto3" json:"publication_data,omitempty"`
// Tags
