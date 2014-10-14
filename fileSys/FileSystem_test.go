package fileSys

import (
	"testing"
	"reflect"
)

func TestGetFileSys(t *testing.T) {
	for i:=0;i<10;i++{
		f1:=GetFileSys()
		f2:=GetFileSys()
		if !reflect.DeepEqual(f1, f2){
			t.Errorf("different file system")
		}
	}
}