package common

import (
	"net"
	"strconv"
	"log"
	"fmt"
	"packets"
	"timeoutcontroller"

)

type Common struct{
	conn *net.UDPConn 
	timeout	timeoutcontroller.Duration
}
func(com *Common) Close(){
	com.conn.Close()
}

// get the UDPConn instance
func NewServerConnection(host string,port int) *Common{
	bindAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(host, strconv.Itoa(port)))
	if err != nil {
		panic(err.Error())
	}

	udpConn, err := net.ListenUDP("udp", bindAddr)
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Listening on %v\n", udpConn.LocalAddr())
	return &Common{
		conn:udpConn,
		timeout : timeoutcontroller.NewDuration(0),
	}
}

func NewUDPConnection() (*Common,error){
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 0,
	})
	if err != nil {
		log.Println("Error listening", err)
		return nil,nil
	}
	return &Common{
		conn : conn,
		timeout	: timeoutcontroller.NewDuration(1),
		},err
}



func (serverConn *Common) ErrorCheck(err error, n int, remoteAddr net.Addr,packet []byte) (*RequestPacket, error){
	if err != nil {
		return nil,fmt.Errorf("Error reading from connection: %v", err)
	}
	if n > packets.BlockSize {
		return nil,fmt.Errorf("Packet too big: %d bytes", n)
	}

	log.Printf("Request from %v", remoteAddr)
	req, err := ParseRequestPacket(packet)
	if err != nil {
		return req, fmt.Errorf("Error parsing request packet: %v", err)
	}
	if !packets.AcceptedMode(req.Mode) {
		return req,fmt.Errorf("Unknown mode: %s", req.Mode)
	}	
	return req,nil
}
