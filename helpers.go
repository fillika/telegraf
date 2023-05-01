package telegraf

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func createUrlWithTokenAndMethod(token string, method string) string {
	url := telegram_api + token + "/" + method
	return url
}

// this medthod make only POST requests
func makeRequest(url string, jsonStr []byte) (response *http.Response, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// todo decide how to create a common method for all requests
func getUpdates() (*http.Response, error) {
	token, err := os.ReadFile(".env")

	if err != nil {
		return nil, err
	}

	createdUrl := createUrlWithTokenAndMethod(string(token), methods.getUpdates)

	params := url.Values{}
	params.Set("offset", "0")
	params.Set("limit", "100")
	params.Set("timeout", "50")
	params.Set("allowed_updates", "[]")

	response, err := makeRequest(createdUrl, []byte(""))

	if err != nil {
		return nil, err
	}

	return response, nil
}

func decodeResponse(response *http.Response) ([]*Update, error) {
	data, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	result := &Response{}

	err = json.Unmarshal(data, result)

	response.Body.Close()
	return result.Result, err
}
