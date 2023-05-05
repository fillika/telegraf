package telegraf

import (
	"encoding/json"
	"log"
	"time"
)

type BotAPI struct {
	token string
	user  User
}

type UpdatesConfig struct {
	Offset int `json:"offset"`
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

// https://core.telegram.org/bots/api#sendmessage
func (b *BotAPI) SendMessage(config MessageConfig) (Message, error) {
	url := createUrlWithTokenAndMethod(b.token, methods.sendMessage)
	msg, err := config.makeRequest(url)

	return msg, err
}
