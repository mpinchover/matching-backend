package main

import (
	"context"
	"log"
	"matching/src/controllers/matching"
	"matching/src/handlers"
	"matching/src/middleware"
	"matching/src/repo"
	"matching/src/route"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()
	fx.New(
		fx.Provide(repo.New),
		fx.Provide(matching.New),
		fx.Provide(handlers.New),
		fx.Provide(NewGinRouter),
		fx.Invoke(SetupRoutes), // kicks everything off
	).Run()
}

type SetupRoutesParams struct {
	fx.In

	Handler *handlers.Handler
	Router  *gin.Engine
}

func SetupRoutes(p SetupRoutesParams) {

	// auth
	p.Router.GET("/test-auth", middleware.EnsureValidToken, route.NewRoute(p.Handler.TestAuth))

	p.Router.POST("/save-tracked-like", middleware.EnsureValidToken, route.NewRoute(p.Handler.SaveTrackedLike))
}

func NewGinRouter(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Opening server on port 9092")
				err := http.ListenAndServe(":9092", r)
				if err != nil {
					log.Println(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
	return r
}
