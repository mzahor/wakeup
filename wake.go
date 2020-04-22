package wakeup

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

const (
	INVALID_MAC = iota
	UDP_CONN_ERR
	UDP_SEND_ERR
)

var errDescr = map[int]string{
	INVALID_MAC:  "Invalid mac address format",
	UDP_CONN_ERR: "Failed to create UDP connection",
	UDP_SEND_ERR: "Failed to send a magic packet",
}

type WakeupError struct {
	Code          int
	Description   string
	OriginalError error
}

func (e *WakeupError) Error() string {
	if e.OriginalError != nil {
		return fmt.Sprintf("%s, %s", e.Description, e.OriginalError)
	}
	return fmt.Sprintf("%s", e.Description)
}

func newErr(code int, origErr error) *WakeupError {
	return &WakeupError{
		Code:          code,
		Description:   errDescr[code],
		OriginalError: origErr,
	}
}

func WakeUp(ip, mac string) error {
	macAddr, err := hex.DecodeString(strings.ReplaceAll(mac, ":", ""))
	if err != nil || len(macAddr) != 6 {
		return newErr(INVALID_MAC, nil)
	}
	ipAddr := net.ParseIP(ip)
	// we always want to broadcast
	ipAddr[3] = 0xff
	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP: ipAddr, Port: 7})
	if err != nil {
		return newErr(UDP_CONN_ERR, err)
	}

	magic := make([]byte, 102)
	copy(magic[:6], []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	for i := 1; i <= 16; i++ {
		copy(magic[i*6:], macAddr)
	}

	_, err = conn.Write(magic)
	if err != nil {
		return newErr(UDP_SEND_ERR, err)
	}

	return nil
}
