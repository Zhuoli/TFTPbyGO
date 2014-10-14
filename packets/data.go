
package packets
import	(
	"encoding/binary"
	)

const DataOpcode uint16 = 3

//  2 bytes     2 bytes      n bytes
//  ----------------------------------
// | Opcode |   Block #  |   Data     |
//  ----------------------------------
func CreateDataPacket(blockNumber uint16, data []byte) []byte {
	buf := make([]byte, 2+2+len(data))
	binary.BigEndian.PutUint16(buf, uint16(DataOpcode))
	binary.BigEndian.PutUint16(buf[2:], blockNumber)
	copy(buf[4:], data)
	return buf
}