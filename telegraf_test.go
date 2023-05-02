package telegraf

import (
	"os"
	"testing"
)

var botAPI *BotAPI

func TestNewBot(t *testing.T) {
	var err error

	token, _ := os.ReadFile(".env")
	botAPI, err = NewBot(string(token))

	if err != nil {
		t.Error("Test 'TestNewBot' failed")
	}
}

func TestNewMessage(t *testing.T) {
	message := NewMessage(1, "Hello World")

	if message.ChatID != 1 || message.Text != "Hello World" {
		t.Error("Test 'TestNewMessage' failed")
	}

	if message.DisableNotification {
		t.Error("Test 'TestNewMessage' failed")
	}

	message.MessageThreadId = 42
	message.DisableNotification = true

	if message.MessageThreadId != 42 || message.DisableNotification != true {
		t.Error("Test 'TestNewMessage' failed")
	}
}
