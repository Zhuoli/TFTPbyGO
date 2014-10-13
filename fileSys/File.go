package fileSys

import (
	"io"
)

type File struct{
	filename string
	buffer [] byte
}
var EOF = io.EOF


func newFile(filename string)(*File,error){
	fil :=&File{
		filename : filename,
		buffer	 : make([]byte, 0),
	} 
	return fil,nil
}

func (F *File) Size()int{
	return len(F.buffer)
}


