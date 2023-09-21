package matching

import (
	"matching/src/gateways/messaging"
	"matching/src/gateways/storage"
	"matching/src/repo"
	"matching/src/types/entities"
	"matching/src/types/requests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

type MatchingController struct {
	repo      repo.IRepo
	storage   storage.IStorage
	messaging messaging.IMessaging
}

type Params struct {
	fx.In

	Repo      *repo.Repo
	Storage   *storage.Storage
	Messaging *messaging.Messaging
}

func New(p Params) *MatchingController {
	return &MatchingController{
		repo:      p.Repo,
		storage:   p.Storage,
		messaging: p.Messaging,
	}
}

func (m *MatchingController) SaveTrackedQuestion(c *gin.Context, tq *entities.TrackedQuestion) (*requests.SaveTrackedQuestionResponse, error) {
	existingTrackedQuestion, err := m.repo.GetTrackedQuestionByUserAndQuestion(tq.UserUUID, tq.QuestionUUID)
	if err != nil {
		return nil, err
	}
	if existingTrackedQuestion == nil {
		err = m.repo.CreateTrackedQuestion(tq)
		if err != nil {
			return nil, err
		}
	} else {
		err = m.repo.UpdateTrackedQuestion(tq)
		if err != nil {
			return nil, err
		}
	}

	// check only for users who have answered the THRESHOLD amount of questions
	result, err := m.FindCompatibleCandidates(tq)
	if err != nil {
		return nil, err
	}

	resp := &requests.SaveTrackedQuestionResponse{}

	// start with just sending back the matched candidates, later add it to question feed
	if len(result.Candidates) > 0 {
		resp.Candidates = result.Candidates
	}
	return resp, nil
}

func (m *MatchingController) SaveTrackedLike(c *gin.Context, tl *entities.TrackedLike) (*requests.SaveTrackedLikeResponse, error) {
	existingTrackedLike, err := m.repo.GetTrackedLikeByUserAndTarget(tl.UserUUID, tl.TargetUUID)
	if err != nil {
		return nil, err
	}

	if existingTrackedLike == nil {
		err = m.repo.CreateTrackedLike(tl)
		if err != nil {
			return nil, err
		}
	} else {
		err = m.repo.UpdateTrackedLike(existingTrackedLike)
		if err != nil {
			return nil, err
		}
	}

	targetLikesUserTrackedLike, err := m.repo.GetTrackedLikeByUserAndTarget(tl.TargetUUID, tl.UserUUID)
	if err != nil {
		return nil, err
	}

	// target has also liked the user
	if targetLikesUserTrackedLike != nil {
		// create a room
		room := &entities.Room{
			Members: []*entities.Member{
				{UserUUID: targetLikesUserTrackedLike.UserUUID},
				{UserUUID: targetLikesUserTrackedLike.TargetUUID},
			},
		}

		createdRoom, err := m.messaging.CreateRoom(room)
		if err != nil {
			return nil, err
		}

		match := &entities.Match{
			UUID:     uuid.New().String(),
			RoomUUID: createdRoom.UUID,
		}
		err = m.repo.CreateMatch(match)
		if err != nil {
			return nil, err
		}

		return &requests.SaveTrackedLikeResponse{
			Match: match,
		}, nil
	}

	return &requests.SaveTrackedLikeResponse{}, nil
}
