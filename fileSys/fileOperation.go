package fileSys

import (
	"errors"
	"fmt"
)

func Create(filename string, fs *FileSys)(*File,error){
	fs.Lock()
	_,ok:=fs.FileMap[filename]
	defer fs.Unlock()
	if ok{
		return nil,errors.New("File already exists.")
	}
	fs.FileMap[filename]=nil
	return newFile(filename)
}


func (fs *FileSys)Open(filename string) (* File,error){
	fil,ok :=fs.FileMap[filename]
	if !ok || fil==nil{
		err := fmt.Errorf("open %s: no such file or directory", filename)
		return nil,err
	}

	return fil,nil
}
// nothing to do
func (f *File)Close(){
	//sync mutex
//	f.file.Close()
}


