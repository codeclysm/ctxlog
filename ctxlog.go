package ctxlog

import (
	"context"

	"logur.dev/logur"
)

// Logger is an interface very similar to the one from logur.dev,
// except for the first parameter which is context
type Logger interface {
	Trace(ctx context.Context, msg string, fields ...map[string]interface{})
	Debug(ctx context.Context, msg string, fields ...map[string]interface{})
	Info(ctx context.Context, msg string, fields ...map[string]interface{})
	Warn(ctx context.Context, msg string, fields ...map[string]interface{})
	Error(ctx context.Context, msg string, fields ...map[string]interface{})
}

type logKey string

// LogKey is the key used to store and retrieve the entry log
// You'll probably want to store the initial entry at the very beginning of the request/trace/whatever
var LogKey = logKey("ctxlog")

// WithFields adds a map of fields to the context.
// If a map already exists, entries will be merged. Keys with the same name will be overwritten
func WithFields(ctx context.Context, newFieldsList ...map[string]interface{}) context.Context {
	fields := mergeFields(ctx, newFieldsList)

	return context.WithValue(ctx, LogKey, fields[0])
}

// CtxLogger is a wrapper around a regular logur Logger.
// When one of its method is called. it forwards to the logger not only the fields specified
// in the method call, but also fields that were saved into the context with the function WithFields
type CtxLogger struct {
	Log logur.Logger
}

func (c CtxLogger) Trace(ctx context.Context, msg string, fields ...map[string]interface{}) {
	fields = mergeFields(ctx, fields)
	c.Log.Trace(msg, fields...)
}
func (c CtxLogger) Debug(ctx context.Context, msg string, fields ...map[string]interface{}) {
	fields = mergeFields(ctx, fields)
	c.Log.Debug(msg, fields...)
}
func (c CtxLogger) Info(ctx context.Context, msg string, fields ...map[string]interface{}) {
	fields = mergeFields(ctx, fields)
	c.Log.Info(msg, fields...)
}
func (c CtxLogger) Warn(ctx context.Context, msg string, fields ...map[string]interface{}) {
	fields = mergeFields(ctx, fields)
	c.Log.Warn(msg, fields...)
}
func (c CtxLogger) Error(ctx context.Context, msg string, fields ...map[string]interface{}) {
	fields = mergeFields(ctx, fields)
	c.Log.Error(msg, fields...)
}

// New returns a new CtxLogger from a regular logur Logger
func New(log logur.Logger) CtxLogger {
	return CtxLogger{Log: log}
}

func getFields(ctx context.Context) map[string]interface{} {
	fields, ok := ctx.Value(LogKey).(map[string]interface{})
	if !ok {
		fields = map[string]interface{}{}
	}

	return fields
}

func mergeFields(ctx context.Context, newFieldsList []map[string]interface{}) []map[string]interface{} {
	oldFields := getFields(ctx)

	// We need a new map because we don't want to pollute the old one
	fields := map[string]interface{}{}

	for k, v := range oldFields {
		fields[k] = v
	}

	for _, newFields := range newFieldsList {
		for k, v := range newFields {
			fields[k] = v
		}
	}

	return []map[string]interface{}{fields}
}
