package requestAgent

import (
	"net"
)

// different request case for requestHandler
func (a *RequestHandler) getRequest(remoteAddress net.Addr, filename string){
	a.HandleGetRequest(remoteAddress,filename)
}

func (a *RequestHandler) putRequest(remoteAddress net.Addr, filename string){
	a.HandlePutRequest(remoteAddress,filename)
}
func (a *RequestHandler)handleError(remoteAddress net.Addr, filename string) {
	//to be overwrite
}
func (a *RequestHandler)handleInvalidPacket(remoteAddress net.Addr, filename string) {
	//to be overwrite
}

