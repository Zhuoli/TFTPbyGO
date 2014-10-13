package fileSys

import (
	"io"
	"fmt"
	
	"chanmutex"
)

type File struct{
	ChanLock	chanmutex.ChanLock
	filename string
	buffer [] byte
}
var EOF = io.EOF


func NewFile(filename string)(*File,error){
	fil :=&File{
		ChanLock : chanmutex.NewChanLock(),
		filename : filename,
		buffer	 : make([]byte, 0),
	} 
	return fil,nil
}



func (fs *FileSys)Open(filename string) (* File,error){
	fil,ok :=fs.FileMap[filename]
	if !ok{
		err := fmt.Errorf("open %s: no such file or directory", filename)
		return nil,err
	}
	return &fil,nil
}

func (f *File)Close(){
	//sync mutex
//	f.file.Close()
}

