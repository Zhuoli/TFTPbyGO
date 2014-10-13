package fileSys

import (
	"bytes"
)
type Reader struct{
	buffer *bytes.Buffer
//	fileSys *FileSys
}

func NewReader(F *File) *Reader{
	return &Reader{
		buffer : bytes.NewBuffer(F.buffer),
	}
}

func (rd *Reader) Read(chunk []byte)(n int, err error){
	n,err=rd.buffer.Read(chunk)
	if err!=nil{
		err=EOF
	}
	return n,err
}
