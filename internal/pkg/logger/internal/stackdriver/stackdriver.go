package stackdriver

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	logKeyServiceContext        = "serviceContext"
	logKeyContextHTTPRequest    = "context.httpRequest"
	logKeyContextUser           = "context.user"
	logKeyContextReportLocation = "context.reportLocation"
)

var logLevelSeverity = map[zapcore.Level]string{
	zapcore.DebugLevel:  "DEBUG",
	zapcore.InfoLevel:   "INFO",
	zapcore.WarnLevel:   "WARNING",
	zapcore.ErrorLevel:  "ERROR",
	zapcore.DPanicLevel: "CRITICAL",
	zapcore.PanicLevel:  "ALERT",
	zapcore.FatalLevel:  "EMERGENCY",
}

func NewLogger(servicename, version string) (*zap.Logger, error) {
	return NewConfig().Build(
		zap.Fields(
			LogServiceContext(
				&ServiceContext{
					Service: servicename,
					Version: version,
				},
			),
		),
	)
}

func NewLoggerDevelopment(servicename, version string) (*zap.Logger, error) {
	return NewConfigDevelopment().Build(
		zap.Fields(
			LogServiceContext(
				&ServiceContext{
					Service: servicename,
					Version: version,
				},
			),
		),
	)
}

func NewConfig() zap.Config {
	return zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "eventTime",
			LevelKey:       "severity",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    EncodeLevel,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewConfigDevelopment() zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "eventTime",
			LevelKey:       "severity",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    EncodeLevel,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func EncodeLevel(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(logLevelSeverity[lv])
}

func LogServiceContext(ctx *ServiceContext) zapcore.Field {
	return zap.Object(logKeyServiceContext, ctx)
}

func LogHTTPRequest(req *HTTPRequest) zapcore.Field {
	return zap.Object(logKeyContextHTTPRequest, req)
}

func LogUser(user string) zapcore.Field {
	return zap.String(logKeyContextUser, user)
}

func LogReportLocation(loc *ReportLocation) zapcore.Field {
	return zap.Object(logKeyContextReportLocation, loc)
}
