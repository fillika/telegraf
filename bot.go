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

func (b *BotAPI) GetUpdatesChannel(config UpdatesConfig) (chan *Update, error) {
	updates := make(chan *Update, 100)

	go func() {
		for {
			response, err := getUpdates(b)

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

			for _, update := range decodedResponse {
				if update.UpdateID >= config.Offset {
					config.Offset = config.Offset + 1
				}
				updates <- update
			}
		}
	}()

	return updates, nil
}

func (b *BotAPI) SendMessage(config MessageConfig) (Message, error) {
	createdUrl := createUrlWithTokenAndMethod(b.token, methods.sendMessage)

	bytes, err := json.Marshal(config)

	if err != nil {
		return Message{}, err
	}

	response, err := makeRequest(createdUrl, bytes)

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