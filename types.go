package telegraf

import "encoding/json"

type ApiResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
}

type ApiError struct {
	Code    int
	Message string
}

// Error message string.
func (e ApiError) Error() string {
	return e.Message
}

// https://core.telegram.org/bots/api#sendmessage
type MessageConfig struct {
	ChatID                   int             `json:"chat_id"`
	MessageThreadId          int             `json:"message_thread_id,omitempty"`
	Text                     string          `json:"text"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool            `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageID         int             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup interface{} // todo create interface for markups
}

// Update represents an incoming update.
// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID          int     `json:"update_id"`
	Message           Message `json:"message"`
	EditedMessage     Message `json:"edited_message"`
	ChannelPost       Message `json:"channel_post"`
	EditedChannelPost Message `json:"edited_channel_post"`
	// InlineQuery  InlineQuery `json:"inline_query"`
	// ChosenInlineResult  ChosenInlineResult `json:"chosen_inline_result"`
	// CallbackQuery  	CallbackQuery	 `json:"callback_query"`
	// ShippingQuery  	ShippingQuery	 `json:"shipping_query"`
	// PreCheckoutQuery PreCheckoutQuery	 `json:"pre_checkout_query"`
	// Poll Poll	 `json:"poll"`
	// PollAnswer PollAnswer	 `json:"poll_answer"`
	// PollAnswer PollAnswer	 `json:"poll_answer"`
	// 	ChatMember 	ChatMemberUpdated	 `json:"chat_member"`
	// 	ChatMemberUpdated 	ChatMemberUpdated	 `json:"my_chat_member"`
	// 	ChatJoinRequest 	ChatJoinRequest	 `json:"chat_join_request"`
}

type Message struct {
	// todo create all fields from documentation
	MessageID       int    `json:"message_id"`
	MessageThreadID int    `json:"message_thread_id"`
	From            User   `json:"from"`
	Chat            Chat   `json:"chat"`
	SenderChat      Chat   `json:"sender_chat"`
	Date            int    `json:"date"`
	Text            string `json:"text"`
}

// todo create all fields from documentation
type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// todo create all fields from documentation
type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

// todo
type MessageEntity struct {
}
