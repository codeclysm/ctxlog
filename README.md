ctxlog
======

Ctxlog is an utility for logrus to store an entry in the context

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
    ctxlog.Debug(ctx, "I'm doing stuff")
}

func handler() {
    ctx := ctxlog.WithField(context.Background(), "reqID", "uniqueReqID")
    MyFunction(ctx)
}
```

``` 
DEBU[0001] I'm doing stuff      reqID=uniqueReqID
```

## How does it work

`ctxlog.WithField()` populates a `logrus.Entry` in the given context with the key `ctxlog.LogKey`
`Debug|Info|Warn|Error()` retrieve the `logrus.Entry` from the given context and use it to call the `logrus.Entry.Debug|Info|Warn|Error` function

## FAQ

*Why should I save a logrus.Entry in the context instead of just logging all the context values?*

Because you probably don't want to log everything in the context, there could be structs and pointer in there. Also it's not trivial to walk through a context values

*Why should I pass a `context` if I could just pass an already populated `logrus.Entry`?*

There's no real advantage, except that a lot of functions already take a `context` for other reasons.
