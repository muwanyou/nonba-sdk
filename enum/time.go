package enum

type Timeline string

var (
	TimelinePast    Timeline = "past"
	TimelinePresent Timeline = "present"
	TimelineFuture  Timeline = "future"
)
