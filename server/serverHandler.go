package server

import (
	"log"
)
// while loop, start listen on UDP port
func(s *Server)  Run(){
	for {
		err := s.RequestHandler.Listen(s.ServerConn)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}


