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

func CreateNewBot(token string) (*BotAPI, error) {
	bot, err := getMe(token)

	if err != nil {
		return &BotAPI{}, err
	}

	return bot, nil
}
