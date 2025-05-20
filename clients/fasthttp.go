package clients

import "github.com/valyala/fasthttp"

type FastHttpClient struct {
	Client *fasthttp.Client
	URL    string
}

func NewFastHttpClient(url string) *FastHttpClient {
	return &FastHttpClient{
		Client: &fasthttp.Client{
			MaxConnsPerHost: 10000,
		},
		URL: url,
	}
}

func (f *FastHttpClient) MakeGetRequest() error {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(f.URL)
	req.Header.SetMethod("GET")

	return f.Client.Do(req, resp)
}
