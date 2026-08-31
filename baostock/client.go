package baostock

import (
	"fmt"
	"time"

	"github.com/example/baostock-go/internal/protocol"
	"github.com/example/baostock-go/internal/result"
)

type Client struct {
	serverAddr string
	timeout    time.Duration
	userID     string
	password   string
	apiKey     string

	conn *protocol.Conn
}

func NewClient(opts ...Option) *Client {
	c := &Client{
		serverAddr: "public-api.baostock.com:10030",
		timeout:    30 * time.Second,
		userID:     "anonymous",
		password:   "123456",
	}
	for _, o := range opts {
		o(c)
	}
	return c
}

func (c *Client) Login() error {
	conn := protocol.NewConn(c.serverAddr, c.timeout)
	if err := conn.Connect(); err != nil {
		return err
	}
	options := "0"
	if c.apiKey != "" {
		options = c.apiKey
	}
	if _, err := conn.Login(c.userID, c.password, options); err != nil {
		_ = conn.Close()
		return fmt.Errorf("login failed: %w", err)
	}
	c.conn = conn
	return nil
}

func (c *Client) Logout() error {
	if c.conn == nil {
		return nil
	}
	_, err := c.conn.Logout()
	closeErr := c.conn.Close()
	c.conn = nil
	if err != nil {
		return err
	}
	return closeErr
}

func (c *Client) SendQuery(method string, params ...string) (*result.ResultData, error) {
	if c.conn == nil {
		return nil, fmt.Errorf("client not logged in")
	}
	resp, err := c.conn.SendQuery(method, params...)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode != "0" {
		return toResult(resp), fmt.Errorf("baostock error %s: %s", resp.ErrorCode, resp.ErrorMsg)
	}
	return toResult(resp), nil
}

func toResult(resp *protocol.ResponseData) *result.ResultData {
	return &result.ResultData{
		MsgType:       resp.MsgType,
		MsgBodyLength: resp.MsgBodyLength,
		ErrorCode:     resp.ErrorCode,
		ErrorMsg:      resp.ErrorMsg,
		Method:        resp.Method,
		UserID:        resp.UserID,
		CurPageNum:    resp.CurPageNum,
		PerPageCount:  resp.PerPageCount,
		Fields:        resp.Fields,
		Data:          resp.Data,
		Code:          resp.Code,
		CodeName:      resp.CodeName,
		StartDate:     resp.StartDate,
		EndDate:       resp.EndDate,
		Frequency:     resp.Frequency,
		AdjustFlag:    resp.AdjustFlag,
		Year:          resp.Year,
		YearType:      resp.YearType,
		Quarter:       resp.Quarter,
		Day:           resp.Day,
		Date:          resp.Date,
	}
}
