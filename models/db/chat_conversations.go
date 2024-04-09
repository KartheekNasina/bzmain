// chat_conversations.go
package models

type ChatConversation struct {
	ID            string `db:"id"`
	User1ID       string `db:"user1_id"`
	User2ID       string `db:"user2_id"`
	LastMessageID string `db:"last_message_id"`
}
