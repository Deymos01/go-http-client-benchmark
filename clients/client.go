package clients

type Client interface {
	MakeGetRequest() error
}
