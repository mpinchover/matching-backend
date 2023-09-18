package matching

import (
	"matching/src/gateways/storage"
	"matching/src/repo"
	"matching/src/types/entities"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type MatchingController struct {
	repo    repo.IRepo
	storage storage.IStorage
}

type Params struct {
	fx.In

	Repo    *repo.Repo
	Storage *storage.Storage
}

func New(p Params) *MatchingController {
	return &MatchingController{
		repo:    p.Repo,
		storage: p.Storage,
	}
}

func (m *MatchingController) SaveTrackedQuestion(c *gin.Context, tq *entities.TrackedQuestion) error {
	existingTrackedQuestion, err := m.repo.GetTrackedQuestionByUserAndQuestion(tq.UserUUID, tq.QuestionUUID)
	if err != nil {
		return err
	}
	if existingTrackedQuestion == nil {
		err = m.repo.CreateTrackedQuestion(tq)
		if err != nil {
			return err
		}
	} else {
		err = m.repo.UpdateTrackedQuestion(tq)
		if err != nil {
			return err
		}
	}
	// now kick off the process of matching by questions here.
	// check only for users who have answered the THRESHOLD amount of questions
	go func(){
		
	}()
	return nil
}

func (m *MatchingController) SaveTrackedLike(c *gin.Context, tl *entities.TrackedLike) error {
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
