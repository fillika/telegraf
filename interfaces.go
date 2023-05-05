package telegraf

type Config interface {
	prepareParams() ([]byte, error)
	makeRequest(url string) (interface{}, error)
}
