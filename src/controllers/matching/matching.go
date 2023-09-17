package matching

import (
	"matching/src/repo"
	"matching/src/types/entities"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MatchingController struct {
	repo repo.IRepo
}

type Params struct {
	fx.In

	repo *repo.Repo
}

func New(p Params) *MatchingController {
	return &MatchingController{
		repo: p.repo,
	}
}

func (m *MatchingController) SaveTrackedLike(ctx *gin.Context, tl *entities.TrackedLike) error {
	existingTrackedLike, err := m.repo.GetTrackedLikeByUserAndTarget(tl.UserUUID, tl.TargetUUID)
	if err != nil {
		return err
	}

	if existingTrackedLike == nil {
		err = m.repo.CreateTrackedLike(tl)
		if err != nil {
			return err
		}
	} else {
		err = m.repo.UpdateTrackedLike(existingTrackedLike)
		if err != nil {
			return err
		}

	}
	return nil
}
