package counters

type alterCounter int

func New(value int) alterCounter {
	return alterCounter(value)
}
