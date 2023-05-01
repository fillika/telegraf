package telegraf

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateNewBot(t *testing.T) {
	token, _ := os.ReadFile(".env")

	_, err := CreateNewBot(string(token))

	if err != nil {
		t.Error("Test 'TestCreateNewBot' failed")
	}
}

func Test_getUpdates(t *testing.T) {
	token, _ := os.ReadFile(".env")

	bot, _ := CreateNewBot(string(token))

	if bot == nil {
		t.Error("Test failed. Bot is nil")
	}

	response, err := getUpdates(bot)

	if err != nil {
		t.Error("Test failed. Error: ", err)
	}

	data, err := decodeResponse(response)

	if err != nil {
		t.Error("Test failed during decode. Error: ", err)
	}

	// todo delete it later
	for _, update := range data {
		fmt.Println("update.UpdateID: ", update.UpdateID)
		fmt.Println("update.Message.Text: ", update.Message.Text)
	}
}
