package baostock

import "time"

type Option func(*Client)

func WithServer(addr string) Option {
	return func(c *Client) {
		c.serverAddr = addr
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}

func WithUser(user string) Option {
	return func(c *Client) {
		c.userID = user
	}
}

func WithPassword(pwd string) Option {
	return func(c *Client) {
		c.password = pwd
	}
}

func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}
