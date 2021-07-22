package ctxlog_test

import (
	"context"
	"fmt"

	"github.com/codeclysm/ctxlog/v2"
)

func ExampleWithFields() {
	ctx := ctxlog.WithFields(context.Background(), map[string]interface{}{"banana": true})
	fields := ctx.Value(ctxlog.LogKey).(map[string]interface{})

	fmt.Println(fields)

	// Output:
	// map[banana:true]
}

func ExampleWithFields_merge() {
	ctx := context.Background()
	fields, _ := ctx.Value(ctxlog.LogKey).(map[string]interface{})
	fmt.Println("background", fields)

	ctx = ctxlog.WithFields(ctx, map[string]interface{}{"banana": true})
	fields = ctx.Value(ctxlog.LogKey).(map[string]interface{})
	fmt.Println("banana", fields)

	ctx = ctxlog.WithFields(ctx, map[string]interface{}{"apple": true})
	fields = ctx.Value(ctxlog.LogKey).(map[string]interface{})
	fmt.Println("apple", fields)

	// Output:
	// background map[]
	// banana map[banana:true]
	// apple map[apple:true banana:true]
}
