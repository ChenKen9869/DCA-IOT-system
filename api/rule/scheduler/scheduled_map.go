package scheduler

import (
	"sync"
	"time"
)

var ScheduledMap map[uint]*time.Timer

var SMLock *sync.Mutex = &sync.Mutex{}
