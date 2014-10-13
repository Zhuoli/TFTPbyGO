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

