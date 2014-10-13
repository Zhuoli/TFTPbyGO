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


// return FileMap with a initial file in it, named: server.dat, seattle.jpg, out.txt
func initFileMap() map[string] *File{
	fileMap:=make(map[string] *File)
//	val:=[]byte("hello, this is server file")
//	val = append(val,byte(0))
//	fileMap["server.dat"]=&File{
//		filename	: "server.dat",
//		buffer		: val,
//	}
//	fileMap["seattle.jpg"]=read2File("seattle.jpg")
//	fileMap["out.txt"]=read2File("out.txt")
	return fileMap
}

//
//func read2File(filename string) *File{
//	
//	file, err := os.Open(filename)
//
//    if err != nil {
//        panic(err.Error())
//    }
//    defer file.Close()
//
//    stats, statsErr := file.Stat()
//    if statsErr != nil {
//        panic(err.Error())
//    }
//
//    var size int64 = stats.Size()
//    bytes := make([]byte, size)
//
//    bufr := bufio.NewReader(file)
//    _,err = bufr.Read(bytes)
//    
//    
//	return &File{
//		filename : filename,
//		buffer : bytes,
//	}
//}