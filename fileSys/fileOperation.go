package fileSys

import (
)

func Create(filename string, fs *FileSys)(*File,error){
	fil :=&File{
		filename : filename,
		buffer	 : make([]byte, 0),
	}
	return fil,nil
}


func FileCleanup(f *File) {
	// sync mutex
//	if err := f.file.Sync(); err != nil {
//		log.Printf("Error syncing %s, %v", f.file.Name(), err)
//	}
//	if err := f.file.Close(); err != nil {
//		log.Printf("Error closing file %s, %v", f.file.Name(), err)
//	}
}
