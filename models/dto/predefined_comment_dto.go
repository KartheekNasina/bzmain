package models

type PredefinedCommentDTO struct {
	Id string `json:"id" db:"id"`
	CommentText string `json:"comment_text" db:"comment_text"`
}
