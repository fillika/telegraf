package telegraf

import (
	"fmt"
	"testing"
)

func Test_createUrlWithTokenAndMethod(t *testing.T) {
	token := "x2222hsf"
	url := createUrlWithTokenAndMethod(token, methods.getMe)

	if url != fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, methods.getMe) {
		t.Error("Incorrect url")
	}
}

func Test_getUpdates(t *testing.T) {
	config := UpdatesConfig{Offset: 0}

	_, err := getUpdates(botApi, config)

	if err != nil {
		t.Error("Error while getting updates")
	}
}
