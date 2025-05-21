package enum

type Timeline string

var (
	TimelinePast    Timeline = "past"
	TimelinePresent Timeline = "present"
	TimelineFuture  Timeline = "future"
)

func (t Timeline) String() string {
	return string(t)
}

type TimeUnit string

const (
	TimeUnitYear   TimeUnit = "year"
	TimeUnitMonth  TimeUnit = "month"
	TimeUnitDay    TimeUnit = "day"
	TimeUnitHour   TimeUnit = "hour"
	TimeUnitMinute TimeUnit = "minute"
	TimeUnitSecond TimeUnit = "second"
)

const DefaultTimeUnit = TimeUnitYear

func (t TimeUnit) String() string {
	return string(t)
}
