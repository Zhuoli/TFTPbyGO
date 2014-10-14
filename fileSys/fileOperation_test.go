package fileSys

import (
	"testing"
	"reflect"
)

func TestCreate(t *testing.T) {
	fs :=GetFileSys()
	fs.FileMap["file1"]=nil
	fs.FileMap["file3"]=&File{
		filename:	"file3",
		buffer	 : make([]byte, 0),
		
	}
	testCases :=[]struct{
		filename	string
		expectedFile	*File
	}{
		// file does not exists
		{
			filename : "file2",
			expectedFile	: &File{
				filename: "file2", 
				buffer	 : make([]byte, 0),
				},
		},
		// file in writting
		{
			filename : "file1",
			expectedFile:	nil,
		},
		// file already exists
		{
			filename : "file3",
			expectedFile:	nil,
		},
	}
		for _, tc := range testCases {
			file,_ :=Create(tc.filename,fs)
			if file!=tc.expectedFile &&  !reflect.DeepEqual(tc.expectedFile, file){
				t.Errorf("unexpected file",tc.filename)
			}
		}
	
}

func TestOpen(t *testing.T){
	fs :=GetFileSys()
	fs.FileMap["file1"]=nil
	fs.FileMap["file3"]=&File{
		filename:	"file3",
		buffer	 : make([]byte, 0),
		
	}
	testcase :=[]struct{
		filename	string
		file	*File
	}{
		{
			filename	: "file1",
			file:	fs.FileMap["file1"],
		},
		{
			filename:	"file2",
			file:	nil,
		},
		{
			filename:	"file3",
			file:	fs.FileMap["file3"],
		},
	}
	for _,tc :=range testcase{
		file,_:=fs.Open(tc.filename)
		if file!=tc.file{
			t.Errorf("unexpected file open",tc.filename)
		}
	}
}