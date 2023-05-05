package telegraf

type Config interface {
	prepareParams() ([]byte, error)
}
