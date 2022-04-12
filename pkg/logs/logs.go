package logs

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var logLevel = 1

type Context map[string]string

func NewContext() Context {
	return Context{}
}

func (ctx Context) With(key, value string) Context {
	ctx[key] = value
	return ctx
}

func print(l string, ctx Context, format string, a ...interface{}) {
	var b strings.Builder
	b.WriteString("level=")
	b.WriteString(l)

	b.WriteString(" t=")
	b.WriteString(time.Now().Format("15:04:05"))

	for k, v := range ctx {
		b.WriteString(" ")
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(v)
	}

	b.WriteString(" msg=\"")
	b.WriteString(strings.ReplaceAll(fmt.Sprintf(format, a...), "\"", "\\\""))
	b.WriteString("\"")

	fmt.Println(b.String())
}

func Error(ctx Context, format string, a ...interface{}) {
	print("error", ctx, format, a...)
}

func Warn(ctx Context, format string, a ...interface{}) {
	if logLevel > 1 {
		print("warn", ctx, format, a...)
	}
}

func Info(ctx Context, format string, a ...interface{}) {
	if logLevel > 2 {
		print("info", ctx, format, a...)
	}
}

func Debug(ctx Context, format string, a ...interface{}) {
	if logLevel > 3 {
		print("debug", ctx, format, a...)
	}
}

func Fatal(ctx Context, format string, a ...interface{}) {
	print("fatal", ctx, format, a...)
	os.Exit(1)
}

func init() {
	l := os.Getenv("LOG_LEVEL")
	switch strings.TrimSpace(l) {
	case "error":
		logLevel = 1
	case "warn":
		logLevel = 2
	case "info":
		logLevel = 3
	case "debug":
		logLevel = 4
	case "": // default
		logLevel = 3
	default:
		Fatal(nil, "unknown LOG_LEVEL '%s'", l)
	}
}
