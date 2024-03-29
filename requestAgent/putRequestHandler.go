package requestAgent

import (
	"common"
	"net"
	"log"
	"fmt"
	"encoding/binary"
	
	"packets"
	"fileSys"
)

func (a *RequestHandler)HandlePutRequest(remoteAddress net.Addr, filename string) {
	log.Println("Handling WRQ")
	com,err :=common.NewUDPConnection()	
	if err!=nil{
		log.Println(err)
		return
	}
	defer com.Close()
	f, err := fileSys.Create(filename,&a.fs)
	if err != nil {
		log.Println("Write Error")
		log.Println(err)
		// TODO: This error should indicate what went wrong
		com.SendError(packets.FileAlreadyExists, err.Error(), remoteAddress)
		return
	}
	defer cleanIfNotFlushed(filename,a.fs)
	
	bw := fileSys.NewWriter(f)

	
	err = writeFileLoop(bw, com, remoteAddress)
	if err != nil {
		log.Println("Error receiving file:", err)
		return
	}
	bw.Flush()
	log.Println("Seccesfully received:", filename)
}

func cleanIfNotFlushed(filename string, fs fileSys.FileSys){
	f,ok:=fs.FileMap[filename]
	if !ok{
		return
	}else if f==nil{
		delete(fs.FileMap,filename)
	}
}

func writeFileLoop(w *fileSys.Writer, com *common.Common, remoteAddress net.Addr) error{
	tid := uint16(0)

	// Acknowledge WRQ
	ack := packets.CreateAckPacket(tid)
	_, err := com.WriteTo(ack, remoteAddress)
	if err != nil {
		log.Println(err)
		return err
	}

	packet := make([]byte, packets.MaxPacketSize)
	for {
		tid++
		n, _, err := writeFile(w, com, remoteAddress, packet, tid)
		if err != nil {
			return err
		}

		if n < 4+packets.BlockSize {
			return nil
		}
	}
}
func writeFile(w *fileSys.Writer, com *common.Common, remoteAddress net.Addr, packet []byte, tid uint16) (int, net.Addr, error) {
	// Read data packet
	n, replyAddr, err := com.ReadFrom(packet)
	if err != nil {
		return n, replyAddr, fmt.Errorf("Error reading packet: %v", err)
	}

	if len(packet) < 2 {
		return n,replyAddr, fmt.Errorf("Error getting opcode: Packet too small to get opcode")
	}
	opcode := uint16(binary.BigEndian.Uint16(packet))
	if opcode > 5 {
		return n,replyAddr, fmt.Errorf("Error getting opcode: Unknown opcode: %d", opcode)
	}

	if opcode != packets.DataOpcode {
		return n, replyAddr, fmt.Errorf("Expected DATA packet, got %v\n", opcode)
	}

	packetTID := binary.BigEndian.Uint16(packet[2:4])
	if packetTID != tid {
		com.SendError(5, "Unknown transfer id", remoteAddress)
		return n, replyAddr, fmt.Errorf("Expected TID %d, got %d\n", tid, packetTID)
	}

	// Write data to disk
	_, err = w.Write(packet[4:n])
	if err != nil {
		return n, replyAddr, fmt.Errorf("Error writing: %v", err)
	}

	ack := packets.CreateAckPacket(tid)
	//com.Conn.SetWriteDeadline(time.Now().Add(timeoutcontroller.Timeout * time.Second))
	_, err = com.WriteTo(ack, replyAddr)
	if err != nil {
		return n, replyAddr, fmt.Errorf("Error writing ACK packet: %v", err)
	}

	return n, replyAddr, nil
}