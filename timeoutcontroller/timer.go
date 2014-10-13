package timeoutcontroller
import(
	"time"
)


type Duration struct{
	millisecond	int
}

func Second(sec int)Duration{
	return Duration{
		millisecond	: sec*1000,
	}
}

func (d *Duration) IsZero()bool{
	return d.millisecond==0
}

func (d *Duration) GetDuration() time.Duration{
	return time.Duration(d.millisecond)*time.Millisecond
}

