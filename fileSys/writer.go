package fileSys

import (
	"bytes"
)
type Writer struct{
	buffer *bytes.Buffer
	f		File
}

func NewWriter(F *File) *Writer{
	return &Writer{
		buffer : bytes.NewBuffer(F.buffer),
		f : *F,
	}
}

func(w *Writer)Write(chunk []byte)(int,error){
	n,err:=w.buffer.Write(chunk)
	return n,err
}
// unlock this file once wrote done
func(w *Writer)Flush(){
	fs :=GetFileSys()
	fs.Lock()
	defer fs.Unlock()
	fs.FileMap[w.f.filename]=&File{
		filename : w.f.filename,
		buffer	:  w.buffer.Bytes(),
	}
	
}