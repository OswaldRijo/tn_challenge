package logger

import (
	"context"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ContextLogger struct {
	zap *zap.SugaredLogger
}

var logger *ContextLogger

func (s *ContextLogger) getContextValues(ctx context.Context, args ...interface{}) []interface{} {
	var spanIDStr string
	var requestIDStr string
	var traceIDStr string
	spanID := ctx.Value("spanId")
	if spanID != nil {
		spanIDStr = spanID.(string)
	}
	requestID := ctx.Value("requestId")
	if requestID != nil {
		requestIDStr = requestID.(string)
	}
	traceID := ctx.Value("traceId")
	if traceID != nil {
		traceIDStr = traceID.(string)
	}
	newAgrs := make([]interface{}, len(args)+3)
	for i, a := range args {
		newAgrs[i] = a
	}
	newAgrs[len(args)] = zap.String("span_id", spanIDStr)
	newAgrs[len(args)+1] = zap.String("request_id", requestIDStr)
	newAgrs[len(args)+2] = zap.String("trace_id", traceIDStr)
	return newAgrs
}

func (s *ContextLogger) getContextKeyValues(ctx context.Context, args ...interface{}) []interface{} {
	spanIDStr := "none"
	requestIDStr := "none"
	traceIDStr := "none"
	spanID := ctx.Value("spanId")
	if spanID != nil {
		spanIDStr = spanID.(string)
	}
	requestID := ctx.Value("requestId")
	if requestID != nil {
		requestIDStr = requestID.(string)
	}
	traceID := ctx.Value("traceId")
	if traceID != nil {
		traceIDStr = traceID.(string)
	}
	newAgrs := make([]interface{}, len(args)+6)
	for i, a := range args {
		newAgrs[i] = a
	}
	newAgrs[len(args)] = "span_id"
	newAgrs[len(args)+1] = spanIDStr
	newAgrs[len(args)+2] = "request_id"
	newAgrs[len(args)+3] = requestIDStr
	newAgrs[len(args)+4] = "trace_id"
	newAgrs[len(args)+5] = traceIDStr
	return newAgrs
}

func (s *ContextLogger) With(ctx context.Context, args ...interface{}) *zap.SugaredLogger {
	return s.zap.With(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Debug(ctx context.Context, args ...interface{}) {
	s.zap.Debug(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Info(ctx context.Context, args ...interface{}) {
	s.zap.Info(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Warn(ctx context.Context, args ...interface{}) {
	s.zap.Warn(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Error(ctx context.Context, args ...interface{}) {
	s.zap.Error(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) DPanic(ctx context.Context, args ...interface{}) {
	s.zap.DPanic(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Panic(ctx context.Context, args ...interface{}) {
	s.zap.Panic(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Fatal(ctx context.Context, args ...interface{}) {
	s.zap.Fatal(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Debugf(ctx context.Context, template string, args ...interface{}) {
	s.zap.Debugf(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Infof(ctx context.Context, template string, args ...interface{}) {
	s.zap.Infof(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Warnf(ctx context.Context, template string, args ...interface{}) {
	s.zap.Warnf(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Errorf(ctx context.Context, template string, args ...interface{}) {
	s.zap.Errorf(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) DPanicf(ctx context.Context, template string, args ...interface{}) {
	s.zap.DPanicf(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Panicf(ctx context.Context, template string, args ...interface{}) {
	s.zap.Panicf(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Fatalf(ctx context.Context, template string, args ...interface{}) {
	s.zap.Fatalf(template, s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Debugw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.Debugw(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) Infow(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.Infow(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) Warnw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.Warnw(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) Errorw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.Errorw(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) DPanicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.DPanicw(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) Panicw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.Panicw(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) Fatalw(ctx context.Context, msg string, keysAndValues ...interface{}) {
	s.zap.Fatalw(msg, s.getContextKeyValues(ctx, keysAndValues...)...)
}

func (s *ContextLogger) Debugln(ctx context.Context, args ...interface{}) {
	s.zap.Debugln(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Infoln(ctx context.Context, args ...interface{}) {
	s.zap.Infoln(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Warnln(ctx context.Context, args ...interface{}) {
	s.zap.Warnln(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Errorln(ctx context.Context, args ...interface{}) {
	s.zap.Errorln(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) DPanicln(ctx context.Context, args ...interface{}) {
	s.zap.DPanicln(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Panicln(ctx context.Context, args ...interface{}) {
	s.zap.Panicln(s.getContextValues(ctx, args...)...)
}

func (s *ContextLogger) Fatalln(ctx context.Context, args ...interface{}) {
	s.zap.Fatalln(s.getContextValues(ctx, args...)...)
}

func GetLog() *ContextLogger {
	if logger == nil {

		logger = GetLogger()
	}

	return logger
}

func GetLogger() *ContextLogger {
	if logger == nil {
		logger = new(ContextLogger)
		z := new(zap.Logger)
		if os.Getenv("ENV") == "development" || os.Getenv("ENV") == "dev" {
			config := zap.NewDevelopmentConfig()

			config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
			z, _ = config.Build(zap.AddCallerSkip(1))
		} else {
			config := zap.NewProductionConfig()
			config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
			z, _ = config.Build(zap.AddCallerSkip(1))
		}
		logger.zap = z.Sugar()
	}

	return logger
}
