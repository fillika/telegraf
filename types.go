package telegraf

type Response struct {
	Ok     bool      `json:"ok"`
	Result []*Update `json:"result"`
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
