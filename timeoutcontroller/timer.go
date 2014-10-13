package timeoutcontroller
import(
//	"time"
)

const Timeout =2

type Duration struct{
	second	int
}

func NewDuration(sec int)Duration{
	return Duration{
		second	: sec,
	}
}

func (d *Duration) IsZero()bool{
	return d.second==0
}

//func NewTimeOut()*TimeOut{
//	return  &TimeOut{
//		t	: make(chan bool,1),
//	}
//}

//func (to *TimeOut) wait(duration time.Duration)bool{
//	select{
//		case <- to.t:
//			return true
//			}
//		case <-tim
//	}
//}


//func (to *TimeOut) wait(duration time.Duration){
//	time.Sleep(duration)
//	to.t<-true
//}