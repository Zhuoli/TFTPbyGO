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
	n, remoteAddr, err := serverConn.Conn.ReadFrom(packet)
	req,err :=serverConn.ErrorCheck(err,n,remoteAddr,packet)
	if err!=nil{
		return err
	}
	switch req.OpCode{
		case packets.OpRRQ:
			go a.readRequest(remoteAddr,req.Filename)
		case packets.OpWRQ:
			go a.writeRequest(remoteAddr,req.Filename)
		case packets.OpERROR:
			go a.handleError(remoteAddr,req.Filename)
		default:
			go a.handleInvalidPacket(remoteAddr,req.Filename)
	}
	return nil
}




