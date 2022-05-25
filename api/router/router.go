package router

import (
	"errors"
	"net/http"
	"time"

	"soccer-manager/api/middlewares"
	"soccer-manager/api/service"
	"soccer-manager/api/utils"
	"soccer-manager/config"
	"soccer-manager/constants"
	"soccer-manager/graph"
	"soccer-manager/graph/generated"
	"soccer-manager/instance"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// skipcq: RVV-A0005
func graphqlHandler(dependencies service.Services, introspectionEnabled bool) gin.HandlerFunc {
	h := handler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					Services: dependencies,
				},
			},
		),
	)
	h.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	h.SetQueryCache(lru.New(1000))

	if introspectionEnabled {
		h.Use(extension.Introspection{})
	}
	h.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Init sets up the route for the REST API
func Init(dependencies service.Services) *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.GinContextToContext())

	// setup cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Accept", "Accept-CH", "Accept-Charset", "Accept-Datetime", "Accept-Encoding", "Accept-Ext", "Accept-Features", "Accept-Language", "Accept-Params", "Accept-Ranges", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Access-Control-Expose-Headers", "Access-Control-Max-Age", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Authorization", "Content-Type"}
	corsConfig.AllowAllOrigins = true

	// setup cors middleware
	router.Use(cors.New(corsConfig))

	// panic recovery
	router.Use(nice.Recovery(func(c *gin.Context, err interface{}) {
		var e error
		if err == nil {
			e = nil
		} else {
			e = err.(error)
		}
		utils.HandleError(c, constants.InternalServerError, e)
	}))

	introspectionEnabled := true
	if config.IsProdEnv() {
		introspectionEnabled = false
	}
	if !config.IsDevEnv() {
		router.Use(middlewares.Airbrake(router, instance.Airbrake()))
	}

	router.NoRoute(func(c *gin.Context) {
		utils.HandleError(c, constants.NotFound, errors.New("not found"))
	})

	pingHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}
	router.GET("/ping", pingHandler)

	router.POST("/query", middlewares.Auth(), middlewares.RequestLogger(), graphqlHandler(dependencies, introspectionEnabled))

	if config.IsDevLikeEnv() {
		router.GET("/", playgroundHandler())
	}

	return router
}
