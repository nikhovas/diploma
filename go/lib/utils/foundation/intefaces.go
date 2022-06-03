package foundation

import (
	"context"
	"sync"
)

type Module interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
	Stop()
}

type FunctionalModule interface {
	Module
	Execute()
}