ctxlog
======

Ctxlog is a wrapper around https://logur.dev/ that uses values stored in `context` to enrich log messages.

Imagine you have a function that needs to log something:

```go
func MyFunction() {
    log.Debug("I'm doing stuff")
}

func handler() {
    MyFunction()
}
```

``` 
DEBU[0001] I'm doing stuff 
```

When you are in production, how do you know which lines of log are related to a certain request?

One way would be to send those info through the chain of functions calls, but it's tedious and prone to errors

Imagine if you could just do this


```go
func MyFunction(ctx context.Context) {
    log.Debug(ctx, "I'm doing stuff")
}

func handler() {
    ctx := context.WithValue(context.Background(), "reqID", "uniqueReqID")
    MyFunction(ctx)
}
```

``` 
DEBU[0001] I'm doing stuff      reqID=uniqueReqID
```

Well, with this package you can, although the actual function calls differ a bit:

```go
func MyFunction(ctx context.Context) {
    log.Debug(ctx, "I'm doing stuff")
}

func handler() {
    ctx := ctxlog.WithFields(context.Background(), map[string]interface{}{"reqID": "uniqueReqID"})
    MyFunction(ctx)
}
```

``` 
DEBU[0001] I'm doing stuff      reqID=uniqueReqID
```

## How does it work

`ctxlog.WithFields()` populates a `map[string]interface{}` in the given context with the key `ctxlog.LogKey`
`Debug|Info|Warn|Error()` retrieve the `map[string]interface{}` from the given context and use it to call the `Debug|Info|Warn|Error` function of the underlying logger

## How to use it

If you have a logur logger, it's simple, you just wrap it. In this example we are using a logrus logur logger (try saying that quickly):

```go
import (
    logrusadapter "logur.dev/adapter/logrus"
    github.com/codeclysm/ctxlog/v2
)

func main() {
    // Create logrus logger
    logrusLog := logrus.New()

    logrusLog.SetOutput(os.Stdout)
    logrusLog.SetFormatter(&logrus.TextFormatter{
        EnvironmentOverrideColors: true,
    })

    // Create logur logger from logrus logger
    logurLog := logrusadapter.New(logrusLog)

    // Create ctxlog logger from logur logger
    ctxLog := ctxlog.New(logurLog)

    // You can use it with the context
    log.Debug(context.Background(), "I'm doing stuff") // Will print "I'm doing stuff"

    // Put fields into the context and log again
    ctx := ctxlog.WithFields(context.Background(), map[string]interface{}{
        "reqID": "uniqueReqID",
        "method": "GET"
    })

    log.Debug(ctx, "I'm doing stuff") // Will print "I'm doing stuff reqID=uniqueRedID method-GET"
}


```

## FAQ

*Why should I save a map[string]interface{} in the context instead of just logging all the context values?*

Because you probably don't want to log everything in the context, there could be structs and pointer in there. Also it's not trivial to walk through context values

*Why not use directly logrus?*

Because logur has a cleaner interface with less methods to override

*What happens if there are no values set?*

No additional fields are added, no panics
