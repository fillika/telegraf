package telegraf

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func createUrlWithTokenAndMethod(token string, method string) string {
	url := telegram_api + token + "/" + method
	return url
}

func getMe(token string) (*BotAPI, error) {
	url := createUrlWithTokenAndMethod(token, methods.getMe)

	apiResponse, err := makeRequest(url, []byte(""))

	if err != nil {
		return &BotAPI{}, err
	}

	var user User
	err = json.Unmarshal(apiResponse.Result, &user)

	if err != nil {
		return &BotAPI{}, err
	}

	return &BotAPI{
		token: token,
		user:  user,
	}, nil
}

// this method make only POST requests
func makeRequest(url string, jsonStr []byte) (response *ApiResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	result, err := decodeResponse(resp)

	if err != nil {
		return nil, err
	}

	if !result.Ok {
		return nil, &ApiError{
			Code:    result.ErrorCode,
			Message: result.Description,
		}
	}

	return result, nil
}

func getUpdates(bot *BotAPI, config UpdatesConfig) (*ApiResponse, error) {
	url := createUrlWithTokenAndMethod(bot.token, methods.getUpdates)

	bytes, err := json.Marshal(config)

	if err != nil {
		return nil, err
	}

	response, err := makeRequest(url, bytes)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func decodeResponse(response *http.Response) (*ApiResponse, error) {
	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := ApiResponse{}

	err = json.Unmarshal(data, &result)

	if err != nil {
		return nil, err
	}

	response.Body.Close()
	return &result, err
}
