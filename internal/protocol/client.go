package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type Conn struct {
	mu       sync.Mutex
	conn     net.Conn
	addr     string
	timeout  time.Duration
	userID   string
	loggedIn bool
}

func NewConn(addr string, timeout time.Duration) *Conn {
	return &Conn{addr: addr, timeout: timeout}
}

func (c *Conn) Connect() error {
	conn, err := net.DialTimeout("tcp", c.addr, c.timeout)
	if err != nil {
		return fmt.Errorf("connect failed: %w", err)
	}
	c.conn = conn
	return nil
}

func (c *Conn) Close() error {
	if c.conn == nil {
		return nil
	}
	err := c.conn.Close()
	c.conn = nil
	c.loggedIn = false
	return err
}

func (c *Conn) Login(userID, password, options string) (*ResponseData, error) {
	body := strings.Join([]string{"login", userID, password, options}, messageSplit)
	msgType := requestTypeByMethod("login")
	head := buildMessageHeader(msgType, len(body))
	headBody := head + body
	raw, err := c.sendReceive(headBody + messageSplit + crc32Checksum([]byte(headBody)))
	if err != nil {
		return nil, err
	}
	resp, err := parseMessage(raw)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode != "0" {
		return resp, fmt.Errorf("login failed: %s %s", resp.ErrorCode, resp.ErrorMsg)
	}
	c.userID = userID
	c.loggedIn = true
	return resp, nil
}

func (c *Conn) Logout() (*ResponseData, error) {
	if !c.loggedIn {
		return nil, errors.New("not logged in")
	}
	body := strings.Join([]string{"logout", c.userID, time.Now().Format("20060102150405")}, messageSplit)
	msgType := requestTypeByMethod("logout")
	head := buildMessageHeader(msgType, len(body))
	headBody := head + body
	raw, err := c.sendReceive(headBody + messageSplit + crc32Checksum([]byte(headBody)))
	if err != nil {
		return nil, err
	}
	resp, err := parseMessage(raw)
	if err != nil {
		return nil, err
	}
	if resp.ErrorCode == "0" {
		c.loggedIn = false
	}
	return resp, nil
}

func (c *Conn) SendQuery(method string, params ...string) (*ResponseData, error) {
	if !c.loggedIn {
		return nil, errors.New("you don't login")
	}
	rawMsg, err := buildMessage(method, c.userID, params...)
	if err != nil {
		return nil, err
	}
	raw, err := c.sendReceive(rawMsg)
	if err != nil {
		return nil, err
	}
	return parseMessage(raw)
}

func (c *Conn) sendReceive(msg string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return "", errors.New("connection is not established")
	}

	if err := c.conn.SetDeadline(time.Now().Add(c.timeout)); err != nil {
		return "", err
	}
	if _, err := c.conn.Write([]byte(msg + "\n")); err != nil {
		return "", fmt.Errorf("send failed: %w", err)
	}

	var buf []byte
	tmp := make([]byte, 8192)
	for {
		n, err := c.conn.Read(tmp)
		if n > 0 {
			buf = append(buf, tmp[:n]...)
			if bytes.HasSuffix(buf, []byte("\n")) {
				break
			}
		}
		if err != nil {
			if len(buf) > 0 {
				break
			}
			return "", fmt.Errorf("receive failed: %w", err)
		}
	}

	head := string(buf[:messageHeaderLength])
	payload := buf[messageHeaderLength:]

	headerParts := strings.Split(head, messageSplit)
	if len(headerParts) >= 3 && isCompressedResponseType(headerParts[1]) {
		return handleCompressedPayload(head, payload)
	}
	return string(buf), nil
}
