package telegraf

import (
	"log"
	"net/http"
)

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

func CreateNewBot(token string) *http.Response {
	url := createUrlWithTokenAndMethod(token, methods.getUpdates)

	response, err := makeRequest(url, []byte(""))

	if err != nil {
		// sure?
		log.Fatal(err)
	}

	return response
}
