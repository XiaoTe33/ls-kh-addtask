package utils

import (
	"strconv"
	"sync"
	"time"
)

var lock = &sync.Mutex{}
var baseTimeStamp, _ = time.Parse("2006/01/02 15:04:05", "2023/01/01 00:00:00")
var name = baseTimeStamp.Unix()
var preTimeStamp int64
var num int64

// GetTaskId 仿照雪花算法生成id
func GetTaskId() string {
	lock.Lock()
	TimeStamp := time.Now().Unix() - baseTimeStamp.Unix()
	if preTimeStamp < TimeStamp {
		preTimeStamp = TimeStamp
		num = 0
	} else {
		num++
	}
	id := TimeStamp<<30 | num
	lock.Unlock()
	str := strconv.FormatInt(id, 10)
	return str
}
