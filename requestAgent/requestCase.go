package requestAgent

import (
	"net"
)

// different request case for requestHandler
func (a *RequestHandler) readRequest(remoteAddress net.Addr, filename string){
	a.HandleReadRequest(remoteAddress,filename)
}

func (a *RequestHandler) writeRequest(remoteAddress net.Addr, filename string){
	a.HandleWriteRequest(remoteAddress,filename)
}
func (a *RequestHandler)handleError(remoteAddress net.Addr, filename string) {
	//to be overwrite
}
func (a *RequestHandler)handleInvalidPacket(remoteAddress net.Addr, filename string) {
	//to be overwrite
}

