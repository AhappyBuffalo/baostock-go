package protocol

import (
	"bytes"
	"compress/zlib"
	"errors"
	"io"
	"strconv"
	"strings"
)

func zlibDecompress(data []byte) ([]byte, error) {
	r, err := zlib.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

func handleCompressedPayload(head string, payload []byte) (string, error) {
	parts := strings.Split(head, messageSplit)
	if len(parts) != 3 {
		return "", errors.New("invalid compressed message header")
	}

	bodyLen, err := strconv.Atoi(parts[2])
	if err != nil {
		return "", err
	}
	if bodyLen > len(payload) {
		return "", errors.New("compressed payload shorter than declared length")
	}

	decompressed, err := zlibDecompress(payload[:bodyLen])
	if err != nil {
		return "", err
	}
	return head + string(decompressed), nil
}
