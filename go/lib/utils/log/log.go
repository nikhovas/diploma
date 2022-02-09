package log

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	log2 "log"
)

func Log(ctx context.Context, format string, a ...interface{}) {
	logVal := fmt.Sprintf(format, a...)
	uuidVal := ctx.Value("uuid")
	log2.Printf("%s %s\n", uuidVal, logVal)
}

func NewContext(parent context.Context) context.Context {
	uuidVal := uuid.New().String()
	Log(parent, "Generating new context: %s", uuidVal)
	return context.WithValue(parent, "uuid", uuidVal)
}

func NewRootContext() context.Context {
	uuidVal := uuid.New().String()
	return context.WithValue(context.Background(), "uuid", uuidVal)
}
