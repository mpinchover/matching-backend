package main

import (
	"context"
	"log"
	"matching/src/controllers/matching"
	"matching/src/handlers"
	"matching/src/route"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {

	fx.New(
		fx.Provide(matching.New),
		fx.Provide(handlers.New),
		fx.Invoke(SetupRoutes), // kicks everything off
	).Run()
}

type SetupRoutesParams struct {
	fx.In

	Handler *handlers.Handler
	Router  *gin.Engine
}

func SetupRoutes(p SetupRoutesParams) {
	p.Router.POST("save-tracked-like", route.NewRoute(p.Handler.SaveTrackedLike))
}

func NewGinRouter(lc fx.Lifecycle) *gin.Engine {
	r := gin.Default()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Opening server on port 9090")
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
