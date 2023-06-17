package model

import (
	"github.com/sirupsen/logrus"
	"ls-kh-addtask/log"
)

var myLog = log.Log

func init() {
	myLog.SetLevel(logrus.TraceLevel)
}
