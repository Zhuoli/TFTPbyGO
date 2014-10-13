package common

import (
	"net"
	"time"
)

type readRetrunStat struct{
	n	int
	addr	net.Addr
	err		error
}
func (com *Common) WriteTo(b []byte, addr net.Addr)(int,error){
	return com.conn.WriteTo(b,addr)
}
//
func (com *Common)ReadFrom(b []byte)(int,net.Addr,error){
	// no timeout if timeout is zero
	if com.timeout.IsZero(){
		return com.conn.ReadFrom(b)
	}
	com.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	return com.conn.ReadFrom(b)
}