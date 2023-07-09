package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"time"
)

// Init uses the default log format configuration, which will be written to the 'logs' folder.
// The logs will be kept for seven days.
func Init() (l *logrus.Logger, err error) {
	writer, err := rotatelogs.New(
		path.Join("logs", "%Y-%m-%d.log"),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		logrus.WithError(err).Error("unable to write logs")
		return nil, err
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:    true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
	})
	logrus.AddHook(lfshook.NewHook(writer, &logrus.TextFormatter{
		FullTimestamp:    true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
		ForceQuote:       true,
	}))
	return logrus.StandardLogger(), err
}

func StandardLogger() *logrus.Logger {
	return logrus.StandardLogger()
}

// ModuleLogger - Provides a logrus.Entry for use by modules
// Includes logrus.Fields
func ModuleLogger(name string) *logrus.Entry {
	return logrus.WithField("module", name)
}

// CurrentModuleLogger provides a logrus.Entry with the current module function name
func CurrentModuleLogger() *logrus.Entry {
	moduleFrame := currentFrame(1)
	return ModuleLogger(moduleFrame.Function)
}

// currentFrame retrieve the current function call frame
func currentFrame(extraSkip int) *runtime.Frame {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2+extraSkip, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return &frame
}
