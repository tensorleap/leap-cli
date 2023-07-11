package log

import "github.com/sirupsen/logrus"

var logger = (*logrus.Logger)(nil)

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

func Println( args ...interface{}) {
	logger.Println(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func init() {
	logger = logrus.New()
}
