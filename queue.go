package queue

import (
"time"
"github.com/nicle-lin/redis"
)

type Queue interface {
	PopTask() (value interface{})
	PushTask(value interface{})
	PopTasks(n int) (value []interface{})
	PushTasks(value []interface{}) (n int)

	SetTaskID()
	GetTaskID()

	SetBlock()
	SetBlockTime()
}

type queue struct {
	block     bool
	blockTime time.Time
	* 
}
