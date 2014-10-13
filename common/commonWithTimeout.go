package common

import (
	"net"
	"time"
)

func (com *Common) WriteTo(b []byte, addr net.Addr)(int,error){
	return com.conn.WriteTo(b,addr)
}
//read with timout
func (com *Common)ReadFrom(b []byte)(int,net.Addr,error){
	// no timeout if timeout is zero
	if com.duration.IsZero(){
		return com.conn.ReadFrom(b)
	}
	com.conn.SetReadDeadline(time.Now().Add(com.duration.GetDuration()))
	return com.conn.ReadFrom(b)
}