# telegraf
Library for telegram API bot

[Official telegram documentation](https://core.telegram.org/bots/api)

> IMPORTANT: 
> This library is under development and is not ready for production use.

The main idea of this library is to provide a simple way to create telegram bots. The library is based on the official telegram API. All methods that are available in the official API are available in this library.

Example for creating a bot and gets updates:

```go
func main() {
    token := "Your token"
    botAPI := telegraf.NewBot(token)

    ch, err := botAPI.GetUpdatesChannel()

    if err != nil {
        // handle error
    }

    for update := range ch {
        msg := telegraf.NewMessage(update.Message.Chat.ID, "Hello world!")

        switch update.Message.Text {
        case "Hello":
            msg.Text = "Hello, " + update.Message.From.FirstName
            // This method is the same as sendMessage in documentation
            bot.SendMessage(msg) 
        case "start":
            msg.Text = "Select option"
            bot.SendMessage(msg)
        }
    }
}
```