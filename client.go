package goanda

import (
	"net/http"
)

type Client struct {
	Token string
	Env string
	client *http.Client
}

func (self *Client) GetBase(url string) string {
	return "https://api-"+self.Env+".oanda.com/v3" + url
}

func (self *Client) Do(req *http.Request) (*http.Response, error) {
	BearerToken := "Bearer " + self.Token
	req.Header.Add("Authorization", BearerToken)
	req.Header.Set("Content-Type", "application/json")
	return self.client.Do(req)
}

func NewFxTradeClient(token string) *Client {
	return &Client{
		Token: token,
		Env: "fxtrade",
		client: &http.Client{},
	}
}

func NewFxPracticeClient(token string) *Client {
	return &Client{
		Token: token,
		Env: "fxpractice",
		client: &http.Client{},
	}
}
