package protocol

import "strings"

func zeroPad(content string, length int, left bool) string {
	if len(content) >= length {
		return content
	}
	if left {
		return strings.Repeat("0", length-len(content)) + content
	}
	return content + strings.Repeat("0", length-len(content))
}
