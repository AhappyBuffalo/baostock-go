package protocol

import (
	"hash/crc32"
	"strconv"
)

func crc32Checksum(data []byte) string {
	return strconv.FormatUint(uint64(crc32.ChecksumIEEE(data)), 10)
}

