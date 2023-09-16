package matching

import (
	"matching/src/types/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MatchingController struct {
}

type Params struct {
	fx.In
}

func New(p Params) *MatchingController {
	return &MatchingController{}
}

func (m *MatchingController) CreateTrackedLike(ctx *gin.Context, req *requests.CreateTrackedLikeRequest) error {
	return nil
}
