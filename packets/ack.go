package packets

import (
	"encoding/binary"
	"fmt"
)



//  2 bytes     2 bytes
//  ---------------------
// | Opcode |   Block #  |
//  ---------------------
func ParseAckPacket(packet []byte) (tid uint16, err error) {
	op, err := GetOpCode(packet)
	if err != nil {
		return 0, fmt.Errorf("Error getting opcode: %v", err)
	}
	if op != OpACK {
		return 0, fmt.Errorf("Expected ACK packet, got OpCode: %d", op)
	}
	tid = binary.BigEndian.Uint16(packet[2:])
	return tid, nil
}


//
//  2 bytes     2 bytes
//  ---------------------
// | Opcode |   Block #  |
//  ---------------------
func CreateAckPacket(tid uint16) []byte {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint16(buf, uint16(OpACK))
	binary.BigEndian.PutUint16(buf[2:], tid)
	return buf
}