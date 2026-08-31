package protocol

import (
	"errors"
	"strconv"
	"strings"
)

func buildMessageHeader(msgType string, bodyLength int) string {
	return strings.Join([]string{
		clientVersion,
		msgType,
		zeroPad(strconv.Itoa(bodyLength), messageHeaderBodyLength, true),
	}, messageSplit)
}

func buildQueryMessageBody(userID string, params ...string) string {
	args := make([]string, 0, len(params)+3)
	args = append(args, userID)
	args = append(args, params...)
	return strings.Join(args, messageSplit)
}

func buildMessage(method string, userID string, params ...string) (string, error) {
	msgType := requestTypeByMethod(method)
	if msgType == "" {
		return "", errors.New("unsupported method: " + method)
	}
	body := method + messageSplit + buildQueryMessageBody(userID, params...)
	head := buildMessageHeader(msgType, len(body))
	headBody := head + body
	return headBody + messageSplit + crc32Checksum([]byte(headBody)), nil
}
