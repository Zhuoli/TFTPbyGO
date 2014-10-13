package server

import (
	"log"
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
	
	"fileSys"
)
// while loop, start listen on UDP port
func(s *Server)  Run(){
	go interaction()
	for {
		err := s.RequestHandler.Listen(s.ServerConn)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func interaction(){
	in :=bufio.NewReader(os.Stdin)
	fs :=fileSys.GetFileSys()
	fmt.Println("...")
	for{
		input, err := in.ReadString('\n')
		if err!=nil{
			continue
		}
		input=strings.Trim(input," ")
		input=strings.Trim(input,"\n")
		fmt.Println("cmd: ",input)
		if input=="ls"{
			fmt.Println("file names:")
			for filename,file :=range fs.FileMap{
				fmt.Print(filename)
				if file!=nil{
					fmt.Println("-"+strconv.Itoa(file.Size()))
				}else{
					fmt.Println("- NIL")
				}
			}
		}
	}
}

