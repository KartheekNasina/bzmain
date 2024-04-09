// chat_messages.go
package models

import (
	"time"
)

type ChatMessage struct {
	ID             string      `db:"id"`
	ConversationID string      `db:"conversation_id"`
	SenderID       string      `db:"sender_id"`
	MessageText    string      `db:"message_text"`
	MessageType    MessageType `db:"message_type"`
	Timestamp      time.Time   `db:"timestamp"`
}

type MessageType string

const (
	MessageTypeText  MessageType = "text"
	MessageTypeImage MessageType = "image"
	MessageTypeVideo MessageType = "video"
	MessageTypeAudio MessageType = "audio"
	MessageTypeFile  MessageType = "file"
)
