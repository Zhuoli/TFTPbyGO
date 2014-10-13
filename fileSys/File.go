package fileSys

import (
	"io"
	"fmt"
)

type File struct{
	filename string
	buffer [] byte
}
var EOF = io.EOF

func (fs *FileSys)Open(filename string) (* File,error){
//	f,err := os.Open(filename)
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

