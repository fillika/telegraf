package telegraf

import (
	"encoding/json"
	"log"
	"time"
)

type BotAPI struct {
	token string
	self  User
}

// To get updates we need to send request with config
// https://core.telegram.org/bots/api#getupdates
type UpdatesConfig struct {
	Offset         int      `json:"offset,omitempty"`
	Limit          int      `json:"limit,omitempty"`
	Timeout        int      `json:"timeout,omitempty"`
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

func (b *BotAPI) GetUpdatesChannel(config UpdatesConfig) (chan *Update, error) {
	updates := make(chan *Update, 100)

	go func() {
		for {
			response, err := getUpdates(b, config)

			if err != nil {
				log.Println(err)
				log.Println("Getting update has failed, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)
				continue
			}

			var decodedResponse []*Update
			err = json.Unmarshal(response.Result, &decodedResponse)

			if err != nil {
				log.Println(err)
				panic(err)
			}

			// As I can understand fonr the documentation I send offset as a parameter
			// and Telegram API decide which updates to send me
			// more detail here https://core.telegram.org/bots/api#getupdates
			for _, update := range decodedResponse {
				if update.UpdateID >= config.Offset {
					config.Offset = update.UpdateID + 1
					updates <- update
				}
			}
		}
	}()

	return updates, nil
}

func (b *BotAPI) GetMe() (*User, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.getMe)

	apiResponse, err := makeRequest(url, []byte(""))

	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(apiResponse.Result, &user)

	if err != nil {
		return nil, err
	}

	b.self = user

	return &user, nil
}

// https://core.telegram.org/bots/api#copymessage
func (b *BotAPI) CopyMessage(config CopyMessageConfig) (MessageID, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.copyMessage)
	msgID, err := config.makeRequest(url)

	return msgID, err
}

// https://core.telegram.org/bots/api#forwardmessage
func (b *BotAPI) ForwardMessage(config ForwardMessageConfig) (Message, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.forwardMessage)
	msg, err := config.makeRequest(url)

	return msg, err
}

// TODO: create 2 tests for methods LogOut and Close
// https://core.telegram.org/bots/api#logout
func (b *BotAPI) LogOut() (*ApiResponse, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.logOut)
	response, err := makeRequest(url, []byte(""))

	return response, err
}

// https://core.telegram.org/bots/api#close
func (b *BotAPI) Close() (*ApiResponse, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.close)
	response, err := makeRequest(url, []byte(""))

	return response, err
}

// https://core.telegram.org/bots/api#sendmessage
func (b *BotAPI) SendMessage(config MessageConfig) (Message, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.sendMessage)
	msg, err := config.makeRequest(url)

	return msg, err
}
