// Package ctxlog allows to share logrus entries through the context in a simple way
//
// - LogKey allows to set and retrieve the logrus entry to and from the context.
//
// - WithField and WithFields allows to set fields on the logrus entry in a safe way
//
// - Debug, Info, Warn and Error allows to launch those functions on the logrus entry in a safe way
//
// If there is no logrus entry defined it will return the default one, so it will never panic
package ctxlog

import (
	"context"

	"github.com/sirupsen/logrus"
)

type logKey string

// LogKey is the key used to store and retrieve the entry log
// You'll probably want to store the initial entry at the very beginning of the request/trace/whatever
var LogKey = logKey("ctxlog")

// WithField adds a single field to the entry saved in the given context.
// If there is no entry saved in the context, the default one will be created
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	entry := getEntry(ctx)

	entry = entry.WithField(key, value)

	return context.WithValue(ctx, LogKey, entry)
}

// WithFields adds a map of fields to the entry saved in the given context.
// If there is no entry saved in the context, the default one will be created
func WithFields(ctx context.Context, fields logrus.Fields) context.Context {
	entry := getEntry(ctx)

	entry = entry.WithFields(fields)

	return context.WithValue(ctx, LogKey, entry)
}

// Debug will call the Debug function on the entry saved in the given context.
// If there is no entry saved in the context, the default one will be created
func Debug(ctx context.Context, args ...interface{}) {
	entry := getEntry(ctx)

	entry.Debug(args...)
}

// Info will call the Info function on the entry saved in the given context.
// If there is no entry saved in the context, the default one will be created.
// See the Debug Example
func Info(ctx context.Context, args ...interface{}) {
	entry := getEntry(ctx)

	entry.Info(args...)
}

// Warn will call the Warn function on the entry saved in the given context.
// If there is no entry saved in the context, the default one will be created.
// See the Debug Example
func Warn(ctx context.Context, args ...interface{}) {
	entry := getEntry(ctx)

	entry.Warn(args...)
}

// Error will call the Warn function on the entry saved in the given context.
// If there is no entry saved in the context, the default one will be created.
// See the Debug Example
func Error(ctx context.Context, args ...interface{}) {
	entry := getEntry(ctx)

	entry.Error(args...)
}

func getEntry(ctx context.Context) *logrus.Entry {
	entry, ok := ctx.Value(LogKey).(*logrus.Entry)
	if !ok {
		entry = logrus.NewEntry(logrus.New())
	}

	return entry
}
