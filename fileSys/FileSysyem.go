package fileSys
import (
	"log"
	"io"
	"os"
	"bufio"
	"bytes"
	
	"chanmutex"
	
)

type FileSys struct{
	// key is the file name
	// value is the encapsulated file data
	FileMap map[string] File
}

var fileSystemSingleton *FileSys=nil

func GetFileSys() *FileSys {
	if fileSystemSingleton==nil{
		fileSystemSingleton = &FileSys{
			FileMap : initFileMap(),
		}
	}
		
	return fileSystemSingleton
}

func (fs *FileSys) ReadFileByName(filename string) (*Reader,*File,error){
	f, err := fs.Open(filename)	
	if err != nil {
		log.Println(err)
		return nil,nil,err
	}
	br := NewReader(f)
	return br,f,err
}

func (fs *FileSys) Register(filename string,F File)bool{
	fs.FileMap[filename]=F
	return true
}

// return FileMap with a initial file in it, named: server.dat, seattle.jpg, out.txt
func initFileMap() map[string] File{
	fileMap:=make(map[string] File)
	val:=[]byte("hello, this is server file")
	val = append(val,byte(0))
	fileMap["server.dat"]=File{
		ChanLock : chanmutex.NewChanLock(),
		filename	: "server.dat",
		buffer		: val,
	}
	fileMap["seattle.jpg"]=read2File("seattle.jpg")
	fileMap["out.txt"]=read2File("out.txt")
	return fileMap
}

func read2File(filename string) File{
	file,err :=os.Open(filename)
	if err!=nil{
		panic(err.Error())
	}
	defer file.Close()
	
	reader :=bufio.NewReader(file)
	buffer :=bytes.NewBuffer(make([]byte,0))
	
	chunk := make([]byte, 512)
	for{
		n,err :=reader.Read(chunk)
		if err!=nil && err!=io.EOF{
			panic(err)
		}
		if n==0{
			break
		}
		n,err=buffer.Write(chunk[0:n])
	}
	
	return File{
		ChanLock : chanmutex.NewChanLock(),
		filename : filename,
		buffer : buffer.Bytes(),
	}
}