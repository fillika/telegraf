package telegraf

const (
	telegram_api = "https://api.telegram.org/bot"
)

type Methods struct {
	getMe      string
	getUpdates string
}

var methods = Methods{
	getMe:      "getMe",
	getUpdates: "getUpdates",
}

func CreateNewBot(token string) (*BotAPI, error) {
	url := createUrlWithTokenAndMethod(token, methods.getUpdates)

	_, err := makeRequest(url, []byte(""))

	if err != nil {
		return nil, err
	}

	BotAPI := BotAPI{
		token: token,
	}

	return &BotAPI, nil
}
