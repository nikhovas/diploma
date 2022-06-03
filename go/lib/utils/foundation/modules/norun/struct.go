package norun

import (
	"context"
	"sync"
)

type NoRun struct {}

func (nr *NoRun) Run(ctx context.Context, wg *sync.WaitGroup) {}
func (nr *NoRun) Stop() {}


