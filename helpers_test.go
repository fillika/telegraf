package telegraf

import (
	"fmt"
	"testing"
)

func Test_createUrlWithTokenAndMethod(t *testing.T) {
	token := "x2222hsf"
	url := createUrlWithTokenAndMethod(token, methods.getMe)

	if url != fmt.Sprintf("https://api.telegram.org/bot%s/%s", token, methods.getMe) {
		t.Error("Incorrect url")
	}
}

func Test_makeRequest(t *testing.T) {
	token := "x2222hsf"
	url := createUrlWithTokenAndMethod(token, methods.getMe)

	_, err := makeRequest(url, []byte(""))

	if err != nil {
		t.Error("Error while making request")
	}
}
