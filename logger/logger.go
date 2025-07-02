package logger

import (
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"io"
)

type EchoLogrusLogger struct {
	Logger *logrus.Logger
}

func New() *EchoLogrusLogger {
	logrusLogger := logrus.New()
	logrusLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return &EchoLogrusLogger{Logger: logrusLogger}
}

func (l *EchoLogrusLogger) Output() io.Writer {
	return l.Logger.Out
}

func (l *EchoLogrusLogger) SetOutput(w io.Writer) {
	l.Logger.SetOutput(w)
}

func (l *EchoLogrusLogger) SetPrefix(p string) {
	// optional: implement if needed
}

func (l *EchoLogrusLogger) Prefix() string {
	return ""
}

func (l *EchoLogrusLogger) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	default:
		return log.OFF
	}
}

func (l *EchoLogrusLogger) SetLevel(v log.Lvl) {
	switch v {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	default:
		l.Logger.SetLevel(logrus.PanicLevel)
	}
}

func (l *EchoLogrusLogger) Printj(j log.JSON) { l.Logger.WithFields(logrus.Fields(j)).Print() }
func (l *EchoLogrusLogger) Debugj(j log.JSON) { l.Logger.WithFields(logrus.Fields(j)).Debug() }
func (l *EchoLogrusLogger) Infoj(j log.JSON)  { l.Logger.WithFields(logrus.Fields(j)).Info() }
func (l *EchoLogrusLogger) Warnj(j log.JSON)  { l.Logger.WithFields(logrus.Fields(j)).Warn() }
func (l *EchoLogrusLogger) Errorj(j log.JSON) { l.Logger.WithFields(logrus.Fields(j)).Error() }
func (l *EchoLogrusLogger) Fatalj(j log.JSON) { l.Logger.WithFields(logrus.Fields(j)).Fatal() }
func (l *EchoLogrusLogger) Panicj(j log.JSON) { l.Logger.WithFields(logrus.Fields(j)).Panic() }
