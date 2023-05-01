package telegraf

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateNewBot(t *testing.T) {
	token, _ := os.ReadFile(".env")

	res := CreateNewBot(string(token))

	if res.StatusCode != 200 {
		t.Error("Test failed. Status code isn't 200: ", res.StatusCode)
	}

	// todo delete it later
	fmt.Println("response: ", res)
}

func Test_getUpdates(t *testing.T) {
	response, err := getUpdates()

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
