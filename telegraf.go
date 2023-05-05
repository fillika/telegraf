package telegraf

const (
	telegram_api = "https://api.telegram.org/bot"
)

type Methods struct {
	getMe,
	getUpdates,
	sendMessage string
}

var methods = Methods{
	getMe:       "getMe",
	getUpdates:  "getUpdates",
	sendMessage: "sendMessage",
}

func NewBot(token string) (*BotAPI, error) {
	bot, err := getMe(token)

	if err != nil {
		return &BotAPI{}, err
	}

	return bot, nil
}

// This method create a MessageConfig for method "sendMessage".
// If you want to set more properties you cam mutate this object.
// https://core.telegram.org/bots/api#sendmessage
func NewMessage(chatID int, text string) *MessageConfig {
	return &MessageConfig{
		ChatID: chatID,
		Text:   text,
	}
}
