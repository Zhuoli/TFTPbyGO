package requestAgent

import (
	"net"
	"time"
	"log"
	"fmt"
	"packets"
	
	"fileSys"
	"common"	
)

func (a *RequestHandler) HandleGetRequest(remoteAddress net.Addr, filename string) {
	start := time.Now()
	log.Println("Handling RRQ for", filename)
	com,err :=common.NewUDPConnection()
	if err!=nil{
		return
	}
	defer com.Close()

	br,f,err :=a.fs.GetReader(filename)
	if err!=nil{
		com.SendError(packets.FileNotFound, "File not found", remoteAddress)	
		return
	}
	defer f.Close()
	bytesRead, err := readFileLoop(br, com, remoteAddress, packets.BlockSize)
	
	
	if err != nil {
		log.Println("Error handling read:", err)
	}
	log.Printf("Done sending %s. %d bytes in %v", filename, bytesRead, time.Since(start))
}

// ReadFileLoop will read from r in blockSize chunks, sending each chunk to through conn
// to remoteAddr. After each send it will wait for an ACK packet. It will loop until
// EOF on r.
func readFileLoop(r *fileSys.Reader, com *common.Common, remoteAddr net.Addr, blockSize int) (int, error) {
	var tid uint16
	var bytesRead int
	tid=0
	bytesRead=0
	buffer := make([]byte, blockSize)
	ackBuf := make([]byte, 4)
	for {
		tid++
		n, err := r.Read(buffer)
		if err == fileSys.EOF {
			// We're done
			packet := packets.CreateDataPacket(tid, make([]byte, 0))
			n, err = com.WriteTo(packet, remoteAddr)
			break
		}
		if err != nil {
			return bytesRead, fmt.Errorf("Error reading data: %v", err)
		}
		bytesRead += n

		packet := packets.CreateDataPacket(tid, buffer[:n])
		n, err = com.WriteTo(packet, remoteAddr)
		if err != nil {
			return bytesRead, fmt.Errorf("Error writing data packet: %v", err)
		}

		// Read ack
		//com.Conn.SetReadDeadline(time.Now().Add(timeoutcontroller.Timeout * time.Second))
		i, _, err := com.ReadFrom(ackBuf)
		if err != nil {
			return bytesRead, fmt.Errorf("Error reading ACK packet: %v", err)
		}
		if i != 4 {
			return bytesRead, fmt.Errorf("Expected 4 bytes read for ACK packet, got %d", i)
		}
		ackTid, err := packets.ParseAckPacket(ackBuf)
		if err != nil {
			return bytesRead, fmt.Errorf("Error parsing ACK packet: %v", err)
		}
		if ackTid != tid {
			return bytesRead, fmt.Errorf("ACK tid: %d, does not match expected: %d", ackTid, tid)
		}
	}
	return bytesRead, nil
}