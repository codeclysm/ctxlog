package ctxlog_test

import (
	"context"
	"fmt"
	"os"

	"github.com/codeclysm/ctxlog"
	"github.com/sirupsen/logrus"
)

func ExampleWithField() {
	ctx := ctxlog.WithField(context.Background(), "banana", true)
	entry := ctx.Value(ctxlog.LogKey).(*logrus.Entry)

	fmt.Println(entry.Data)

	// Output:
	// map[banana:true]
}
func ExampleWithFields() {
	ctx := ctxlog.WithFields(context.Background(), logrus.Fields{
		"banana": true,
	})
	entry := ctx.Value(ctxlog.LogKey).(*logrus.Entry)

	fmt.Println(entry.Data)

	// Output:
	// map[banana:true]
}

func ExampleDebug() {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.DebugLevel)
	logger.Formatter = &logrus.TextFormatter{
		DisableTimestamp: true,
	}

	ctx := context.WithValue(context.Background(), ctxlog.LogKey, logrus.NewEntry(logger))

	ctx = ctxlog.WithField(ctx, "banana", true)

	ctxlog.Debug(ctx, "message")

	// Output:
	// level=debug msg=message banana=true
}
