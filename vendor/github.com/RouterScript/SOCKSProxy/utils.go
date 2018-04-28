package proxy

import (
	"strconv"

	"encoding/binary"
)

func toPort(port string) (buffer []byte) {
	buffer = make([]byte, 2)
	if port, err := strconv.ParseUint(port, 10, 16); err == nil {
		binary.BigEndian.PutUint16(buffer, uint16(port))
		return
	}
	return nil
}
