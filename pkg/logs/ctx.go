package logs

import "context"

type ctxKeyType uint

const ctxKey ctxKeyType = iota + 1

func CtxWithLogger(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, ctxKey, logger)
}

func LoggerFromCtx(ctx context.Context) Logger {
	loggerAny := ctx.Value(ctxKey)
	if l, ok := loggerAny.(Logger); ok {
		return l
	}

	return NewLogger(LevelInfo)
}
