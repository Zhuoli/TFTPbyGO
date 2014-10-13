package server

import (
	"requestAgent"
	"common"
)
type Server struct{
	ServerConn *common.Common
	RequestHandler *requestAgent.RequestHandler
}
func NewServer(host string, port int) *Server{
	return &Server{
		ServerConn : common.NewServerConnection(host,port),
		RequestHandler : requestAgent.NewRequestHandler(),
	}
}






