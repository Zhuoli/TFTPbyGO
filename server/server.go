package server

import (
	"requestAgent"
	"common"
)
type Server struct{
	ServerConn *common.Common
	RequestHandler *requestAgent.RequestHandler
}

var serverSingleton *Server =nil
func NewServer(host string, port int) *Server{
	if serverSingleton==nil{
		serverSingleton= &Server{
			ServerConn : common.NewServerConnection(host,port),
			RequestHandler : requestAgent.NewRequestHandler(),
		}
	}
	return serverSingleton
}






