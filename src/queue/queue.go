package queue

import (
	"github.com/3stadt/GoTBot/src/structs"
)

var JobQueue = make(map[string](chan structs.Job))

func NewQueue(name string) {
	JobQueue[name] = make(chan structs.Job)
}

func AddJob(queue string, job structs.Job) {
	JobQueue[queue] <- job
}