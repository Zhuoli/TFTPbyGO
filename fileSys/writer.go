package fileSys

import (
	"bytes"
	"chanmutex"
)
type Writer struct{
	buffer *bytes.Buffer
	filename string
}

func NewWriter(F *File) *Writer{
	return &Writer{
		buffer : bytes.NewBuffer(F.buffer),
		filename : F.filename,
	}
}

func(w *Writer)Write(chunk []byte)(int,error){
	n,err:=w.buffer.Write(chunk)
	return n,err
}

func(w *Writer)Flush(){
	fs :=GetFileSys()
	fs.Register(w.filename,File{
			ChanLock : chanmutex.NewChanLock(),
			filename : w.filename,
			buffer	 : w.buffer.Bytes(),
			})
}