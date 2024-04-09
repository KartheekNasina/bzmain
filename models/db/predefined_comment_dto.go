// predefined_comments_dto.go
package models

type PredefinedComment struct {
	ID          string `db:"id"`
	CommentText string `db:"comment_text"`
}
