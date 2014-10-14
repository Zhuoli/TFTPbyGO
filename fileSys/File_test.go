package fileSys

import (
	"testing"
)

func TestSize(t *testing.T) {
	testcase :=[]struct{
		filename	string
		file		*File
		size		int
	}{
		{
			filename:	"file1KB",
			file:		&File{
				filename: "file1KB",
				buffer	 : make([]byte, 1024),
			},
			size:	1024,
		},
		{
			filename:	"file2zero",
			file:	&File{
				filename: "file2zero",
				buffer	 : make([]byte, 0),
			},
		},
	}
	for _, tc := range testcase {
		sz :=tc.file.Size()
		if sz!=tc.size{
			t.Errorf("unexpected file size",tc.filename)
		}
	}
	
}