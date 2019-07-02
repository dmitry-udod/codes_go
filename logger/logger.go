package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func InitLogger() {
	file, err := os.OpenFile("logs/codes.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Fatal("Failed to log to file, using default stderr")
	}
}