package common

import (
	"net"
	"time"
	"errors"
//	"log"
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
	timeout :=make(chan bool,1)
	var ch chan readRetrunStat
	go com.readfrom(b,ch)
	go func(){
		//wait 1 second
		time.Sleep(1e9)
		timeout <- true
	}()
	
	select{
		case state := <-ch:
			return state.n,state.addr,state.err
		case <-timeout:
			return 0,nil,errors.New("TimeOut")
	}
}

func (com *Common) readfrom(b []byte, stat chan readRetrunStat){
	num,adrs,er:=com.conn.ReadFrom(b)
	stat <- readRetrunStat{
		n	: num,
		addr	: adrs,
		err		: er,
	}
}