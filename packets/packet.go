package packets

import (
	"encoding/binary"
	"fmt"
	"strings"
)
type OpCode uint16
const (
	BlockSize     = 512
	MaxPacketSize = BlockSize * 2
)
const (
	OpRRQ   OpCode = 1
	OpWRQ   OpCode = 2
	OpDATA  OpCode = 3
	OpACK   OpCode = 4
	OpERROR OpCode = 5
)

// GetOpCode will attempt to parse the OpCode from the packet passed in
func GetOpCode(packet []byte) (OpCode, error) {
	if len(packet) < 2 {
		return OpERROR, fmt.Errorf("Packet too small to get opcode")
	}
	opcode := OpCode(binary.BigEndian.Uint16(packet))
	if opcode > 5 {
		return OpERROR, fmt.Errorf("Unknown opcode: %d", opcode)
	}
	return opcode, nil
}

func IsAcceptedMode(mode string) bool {
	switch strings.ToLower(mode) {
	case "netascii", "octet", "mail":
		return true
	}
	return false
}