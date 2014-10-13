package fileSys
import (
	"log"
	"sync"
	
)

type FileSys struct{
	// key is the file name
	// value is the encapsulated file data
	sync.RWMutex 
	FileMap map[string] *File
}

var fileSystemSingleton *FileSys=nil

// singleton pattern
func GetFileSys() *FileSys {
	if fileSystemSingleton==nil{
		fileSystemSingleton = &FileSys{
			FileMap : initFileMap(),
		}
	}
		
	return fileSystemSingleton
}


func (fs *FileSys) GetReader(filename string) (*Reader,*File,error){
	f, err := fs.Open(filename)	
	if err != nil {
		log.Println(err)
		return nil,nil,err
	}
	br := NewReader(f)
	return br,f,err
}


func initFileMap() map[string] *File{
	fileMap:=make(map[string] *File)
	return fileMap
}

