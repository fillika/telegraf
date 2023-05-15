package telegraf

const (
	telegram_api = "https://api.telegram.org/bot"
)

type Methods struct {
	close,
	copyMessage,
	forwardMessage,
	getMe,
	getUpdates,
	logOut,
	sendPhoto,
	sendMessage string
}

var methods = Methods{
	close:          "close",
	copyMessage:    "copyMessage",
	forwardMessage: "forwardMessage",
	getMe:          "getMe",
	getUpdates:     "getUpdates",
	logOut:         "logOut",
	sendMessage:    "sendMessage",
	sendPhoto:    "sendPhoto",
}

// This method create a BotAPI object.
func NewBot(token string) (*BotAPI, error) {
	bot := &BotAPI{
		token: token,
	}

	_, err := bot.GetMe()

	if err != nil {
		return &BotAPI{}, err
	}

	return bot, nil
}

// This method create a MessageConfig for method "sendMessage".
// If you want to set more properties you can mutate this object.
// https://core.telegram.org/bots/api#sendmessage
func NewMessage(chatID int, text string) MessageConfig {
	return MessageConfig{
		ChatID: chatID,
		Text:   text,
	}
}

// This method create a ForwardMessageConfig for method "forwardMessage".
// If you want to set more properties you can mutate this object.
// https://core.telegram.org/bots/api#forwardmessage
func NewForwardMessage(chatID, fromChatID, messageID int) ForwardMessageConfig {
	return ForwardMessageConfig{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}

// This method create a CopyMessageConfig for method "copyMessage".
// If you want to set more properties you can mutate this object.
// https://core.telegram.org/bots/api#copymessage
func NewCopyMessage(chatID, fromChatID, messageID int) CopyMessageConfig {
	return CopyMessageConfig{
		ChatID:     chatID,
		FromChatID: fromChatID,
		MessageID:  messageID,
	}
}
