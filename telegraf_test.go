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

	message.MessageThreadID = 42
	message.DisableNotification = true

	if message.MessageThreadID != 42 || message.DisableNotification != true {
		t.Error("Test 'TestNewMessage' failed")
	}
}

func TestNewForwardMessage(t *testing.T) {
	message := NewForwardMessage(1, 2, 3)

	if message.ChatID != 1 || message.FromChatID != 2 || message.MessageID != 3 {
		t.Error("Test 'TestNewForwardMessage' failed")
	}

	if message.MessageThreadID {
		t.Error("Test 'TestNewForwardMessage' failed")
	}

	message.MessageThreadID = true

	if message.MessageThreadID != true {
		t.Error("Test 'TestNewForwardMessage' failed")
	}
}

func TestNewCopyMessage(t *testing.T) {
	message := NewCopyMessage(1, 2, 3)

	if message.ChatID != 1 || message.FromChatID != 2 || message.MessageID != 3 {
		t.Error("Test 'TestNewCopyMessage' failed")
	}

	if message.MessageThreadID {
		t.Error("Test 'TestNewCopyMessage' failed")
	}

	message.MessageThreadID = true

	if message.MessageThreadID != true {
		t.Error("Test 'TestNewCopyMessage' failed")
	}
}
