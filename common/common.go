package common

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	
	"packets"
	
)

type RequestPacket struct {
	OpCode   packets.OpCode
	Filename string
	Mode     string
}



// parses a request packet and package it with RequestPacket struct
//
//  2 bytes     string    1 byte     string   1 byte
// ------------------------------------------------
// | Opcode |  Filename  |   0  |    Mode    |   0  |
// ------------------------------------------------
func ParseRequestPacket(packet []byte) (*RequestPacket, error) {
	// Get opcode
	opcode, err := GetOpCode(packet)
	if err != nil {
		return nil, err
	}
	// Get filename
	reader := bytes.NewBuffer(packet[2:])
	
	//ReadBytes reads until the first occurrence of delim in the input
	filename, err := reader.ReadBytes(byte(0))
	if err != nil {
		return nil, fmt.Errorf("Error reading filename: %v", err)
	}
	// Remove trailing 0
	filename = filename[:len(filename)-1]

	// Get mode
	mode, err := reader.ReadBytes(byte(0))
	if err != nil {
		return nil, fmt.Errorf("Error reading mode: %v", err)
	}
	// Remove trailing 0
	mode = mode[:len(mode)-1]

	return &RequestPacket{
		OpCode:   opcode,
		Mode:     string(mode),
		Filename: string(filename),
	}, nil
}


// GetOpCode will attempt to parse the OpCode from the packet passed in
func GetOpCode(packet []byte) (packets.OpCode, error) {
	if len(packet) < 2 {
		return packets.OpERROR, fmt.Errorf("Packet too small to get opcode")
	}
	opcode := packets.OpCode(binary.BigEndian.Uint16(packet))
	if opcode > 5 {
		return packets.OpERROR, fmt.Errorf("Unknown opcode: %d", opcode)
	}
	return opcode, nil
}

func(com *Common) SendError(code packets.ErrorCode, message string,remoteAddress net.Addr) error {
	conn:=com.conn
	errPacket := packets.CreateErrorPacket(0, message)
	_, err := conn.WriteTo(errPacket, remoteAddress)
	if err != nil {
		return fmt.Errorf("Error writing error packet: %v", err)
	}
	return nil
}

