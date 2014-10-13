package requestAgent

import (
	"common"
	"fileSys"
	"packets"
)

type RequestHandler struct {
	fs	 fileSys.FileSys
}
func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		fs: *fileSys.GetFileSys(),
	}
}


func (a *RequestHandler) Listen(serverConn *common.Common) error {
	packet := make([]byte, packets.MaxPacketSize)
	n, remoteAddr, err := serverConn.ReadFrom(packet)
	req,err :=serverConn.ErrorCheck(err,n,remoteAddr,packet)
	if err!=nil{
		return err
	}
	switch req.OpCode{
		case packets.OpRRQ:
			a.getRequest(remoteAddr,req.Filename)
		case packets.OpWRQ:
			a.putRequest(remoteAddr,req.Filename)
		case packets.OpERROR:
		 	a.handleError(remoteAddr,req.Filename)
		default:
			a.handleInvalidPacket(remoteAddr,req.Filename)
	}
	return nil
}




