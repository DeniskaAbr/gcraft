package gc_enum

type Priority int

const (
	PriorityLow Priority = iota
	PriorityDefault
	PriorityHigh
)
