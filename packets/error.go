package packets
import (
	"encoding/binary"
	"fmt"
)
const ErrorOpcode uint16 = 5

type ErrorCode uint16

const (
	Undefined                    ErrorCode = 0
	FileNotFound                 ErrorCode = 1
	AccessViolation              ErrorCode = 2
	DiskFullOrAllocationExceeded ErrorCode = 3
	IllegalTftpOperation         ErrorCode = 5
	FileAlreadyExists            ErrorCode = 6
	NoSuchUser                   ErrorCode = 7
)

// creates an error packet with the following structure:
//
// 2 bytes     2 bytes      string    1 byte
// -----------------------------------------
// | Opcode |  ErrorCode |   ErrMsg   |   0  |
// -----------------------------------------
func CreateErrorPacket(code ErrorCode, message string) []byte {
	buf := make([]byte, 2+2+len(message)+1)
	
	binary.BigEndian.PutUint16(buf, uint16(ErrorOpcode)) // 2 bytes correct
	
	binary.BigEndian.PutUint16(buf[2:], uint16(code))    // 2 bytes
	
	copy(buf[4:], []byte(message)) // correct
	buf[len(buf)-1] = byte(0)
	fmt.Println(buf)
	return buf
}
