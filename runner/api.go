package runner

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"soccer-manager/api/router"
	"soccer-manager/api/service"
	"soccer-manager/config"

	"soccer-manager/logger"
)

// API is the interface for rest api runner
type API interface {
	Go(ctx context.Context, wg *sync.WaitGroup)
}

type api struct {
}

func (*api) Go(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	logger.Log.Infof("Starting API server on %v...", config.APIPort())
	services := service.Init()

	routerV1 := router.Init(services)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", config.APIPort()),
		Handler:      routerV1,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.ListenAndServe()
}

// NewAPI returns an instance of the REST API runner
func NewAPI() API {
	return &api{}
}
