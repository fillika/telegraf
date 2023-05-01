package telegraf

type BotAPI struct {
	token string
}

func (b *BotAPI) GetUpdatesChannel() (chan *Update, error) {
	updates := make(chan *Update, 100)

	go func() {
		for {
			response, err := getUpdates(b)

			if err != nil {
				panic(err)
			}

			decodedResponse, err := decodeResponse(response)

			if err != nil {
				panic(err)
			}

			for _, update := range decodedResponse {
				updates <- update
			}
		}
	}()

	return updates, nil
}
