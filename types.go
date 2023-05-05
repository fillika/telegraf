package telegraf

import (
	"encoding/json"
)

// ===================
//     API TYPES
// ===================

type ApiResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result,omitempty"`
	ErrorCode   int             `json:"error_code,omitempty"`
	Description string          `json:"description,omitempty"`
}

type ApiError struct {
	Code    int
	Message string
}

// Error message string.
func (e ApiError) Error() string {
	return e.Message
}

// ===================
//    MESSAGES TYPES
// ===================

type CopyMessageConfig struct {
	ChatID                   int             `json:"chat_id"`
	FromChatID               int             `json:"from_chat_id"`
	MessageID                int             `json:"message_id"`
	MessageThreadID          bool            `json:"message_thread_id,omitempty"`
	Caption                  string          `json:"caption,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	CaptionEntities          []MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageID         int             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup		 interface{} `json:"reply_markup,omitempty"` // todo create interface for markups
}

func (cmc *CopyMessageConfig) prepareParams() ([]byte, error) {
	bytes, err := json.Marshal(cmc)
	return bytes, err
}

func (cmc *CopyMessageConfig) makeRequest(url string) (MessageID, error) {
	bytes, err := cmc.prepareParams()

	if err != nil {
		return MessageID{}, err
	}

	response, err := makeRequest(url, bytes)

	if err != nil {
		return MessageID{}, err
	}

	var msgID MessageID
	err = json.Unmarshal(response.Result, &msgID)

	if err != nil {
		return MessageID{}, err
	}

	return msgID, nil
}

// ===================

type ForwardMessageConfig struct {
	ChatID              int  `json:"chat_id"`
	FromChatID          int  `json:"from_chat_id"`
	MessageID           int  `json:"message_id"`
	MessageThreadID     bool `json:"message_thread_id,omitempty"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	ProtectContent      bool `json:"protect_content,omitempty"`
}

func (fmc *ForwardMessageConfig) prepareParams() ([]byte, error) {
	bytes, err := json.Marshal(fmc)
	return bytes, err
}

func (fmc *ForwardMessageConfig) makeRequest(url string) (Message, error) {
	bytes, err := fmc.prepareParams()

	if err != nil {
		return Message{}, err
	}

	response, err := makeRequest(url, bytes)

	if err != nil {
		return Message{}, err
	}

	var msg Message
	err = json.Unmarshal(response.Result, &msg)

	if err != nil {
		return Message{}, err
	}

	return msg, nil
}

// ===================

// https://core.telegram.org/bots/api#sendmessage
type MessageConfig struct {
	ChatID                   int             `json:"chat_id"`
	Text                     string          `json:"text"`
	MessageThreadId          int             `json:"message_thread_id,omitempty"`
	ParseMode                string          `json:"parse_mode,omitempty"`
	Entities                 []MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool            `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool            `json:"disable_notification,omitempty"`
	ProtectContent           bool            `json:"protect_content,omitempty"`
	ReplyToMessageID         int             `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	// ReplyMarkup interface{} // todo create interface for markups
}

func (mc *MessageConfig) prepareParams() ([]byte, error) {
	bytes, err := json.Marshal(mc)
	return bytes, err
}

func (mc *MessageConfig) makeRequest(url string) (Message, error) {
	bytes, err := mc.prepareParams()

	if err != nil {
		return Message{}, err
	}

	response, err := makeRequest(url, bytes)

	if err != nil {
		return Message{}, err
	}

	var msg Message
	err = json.Unmarshal(response.Result, &msg)

	if err != nil {
		return Message{}, err
	}

	return msg, nil
}

// ===================
