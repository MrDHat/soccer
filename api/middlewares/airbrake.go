package middlewares

import (
	"context"
	"sync"

	"github.com/airbrake/gobrake/v4"

	"github.com/gin-gonic/gin"
)

// Airbrake creates a new airbrake middleware
func Airbrake(engine *gin.Engine, notifier *gobrake.Notifier) func(c *gin.Context) {
	return func(c *gin.Context) {
		routeName := routeName(c, engine)
		_, metric := gobrake.NewRouteMetric(context.TODO(), c.Request.Method, routeName)

		c.Next()

		metric.StatusCode = c.Writer.Status()
		_ = notifier.Routes.Notify(context.TODO(), metric)
	}
}

func routeName(c *gin.Context, engine *gin.Engine) string {
	initPathMap(engine)
	route, ok := pathMap[c.HandlerName()]
	if ok {
		return route
	}
	return "UNKNOWN"
}

var (
	pathMapOnce sync.Once
	pathMap     map[string]string
)

func initPathMap(engine *gin.Engine) {
	pathMapOnce.Do(func() {
		pathMap = make(map[string]string)
		for _, ri := range engine.Routes() {
			pathMap[ri.Handler] = ri.Path
		}
	})
}
