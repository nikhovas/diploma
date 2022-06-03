package foundation

import (
	"context"
	"github.com/nikhovas/diploma/go/lib/utils/log"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"sync"
	"time"
)

func GetConfig(config interface{}) {
	filename, _ := filepath.Abs("./config/config.yaml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(yamlFile, config); err != nil {
		panic(err)
	}
}

func RunBlock(module Module) {
	ctx := log.NewRootContext()
	wg := sync.WaitGroup{}

	module.Run(ctx, &wg)

	wg.Wait()
}

func RunPeriodic(ctx context.Context, wg *sync.WaitGroup, worker func(ctx context.Context), interval time.Duration) {
	wg.Add(1)
	go BlockRunPeriodic(ctx, wg, worker, interval)
}

func BlockRunPeriodic(ctx context.Context, wg *sync.WaitGroup, worker func(ctx context.Context), interval time.Duration) {
	defer wg.Done()

	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ctx.Done():
			ticker.Stop()
			break
		case <-ticker.C:
			worker(ctx)
		}
	}
}