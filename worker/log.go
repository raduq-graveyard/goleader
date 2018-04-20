package worker

import (
	"log"
)

//Worker that executes something
type Worker interface {
	Log(string)
}

//WorkerInfo contains metadata of the worker
type workerInfo struct {
	name string
}

//NewLog worker with the specified name
func NewLog() Worker {
	return &workerInfo{
		name: "log",
	}
}

// Execute the command of log some string
func (w workerInfo) Log(m string) {
	log.Printf("%s: %s", w.name, m)
}
