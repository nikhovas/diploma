package main

import (
	"context"
	"gopkg.in/yaml.v3"
	"os"
	"sync"
	"vkProductsUpdater/application"
)

type Runnable interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
}

//func StartProgram(configType reflect.Type, fabric func )

func OpenApplicationConfigFromFile(path string) (application.ApplicationConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return application.ApplicationConfig{}, err
	}
	defer f.Close()

	appConfig := application.ApplicationConfig{}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&appConfig)
	if err != nil {
		return application.ApplicationConfig{}, err
	}

	return appConfig, nil
}

func main() {
	appConfig, err := OpenApplicationConfigFromFile("configs/dev.yaml")
	if err != nil {
		panic(err)
	}

	app, err := application.NewApplication(appConfig)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() { app.Run(ctx, &wg) }()
	wg.Wait()
}
