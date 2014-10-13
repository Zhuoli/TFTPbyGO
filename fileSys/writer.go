package fileSys

import (
	"bytes"
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

func(w *Writer)Flush(fs *FileSys){
	fs.Register(w.filename,File{
			filename : w.filename,
			buffer	 : w.buffer.Bytes(),
			})
}