package handlers

import (
	"matching/src/controllers/discover"
	"matching/src/controllers/matching"
	"matching/src/types/requests"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	MatchingController *matching.MatchingController
	DiscoverController *discover.DiscoverController
}

type Params struct {
	fx.In

	MatchingController *matching.MatchingController
	DiscoverController *discover.DiscoverController
}

func New(p Params) *Handler {
	return &Handler{
		MatchingController: p.MatchingController,
		DiscoverController: p.DiscoverController,
	}
}

// delete the room as well
// do you need to create a match here? Or can you just rely on the room itself?
// func (h *Handler) DeleteMatch(c *gin.Context) (interface{}, error) {
// 	// add validation

// 	req := &requests.DeleteMatchRequest{}

// 	err := c.BindJSON(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, h.MatchingController.DeleteMatch(c, req)
// }

// deprioritize
func (h *Handler) BlockUser(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// add validation

	return nil, nil
}

// func (h *Handler) CreateMatch(c *gin.Context) (interface{}, error) {
// 	// add validation

// 	return nil, nil
// }

func (h *Handler) SaveTrackedLike(c *gin.Context) (interface{}, error) {
	// add validation

	req := &requests.SaveTrackedLikeRequest{}

	err := c.BindJSON(req)
	if err != nil {
		return nil, err
	}

	return nil, h.MatchingController.SaveTrackedLike(c, req.TrackedLike)
}

func (h *Handler) SaveTrackedQuestion(c *gin.Context) (interface{}, error) {
	// validation

	req := &requests.SaveTrackedQuestionRequest{}

	err := c.BindJSON(req)
	if err != nil {
		return nil, err
	}

	return nil, h.MatchingController.SaveTrackedQuestion(c, req.TrackedQuestion)
}

func (h *Handler) GetQuestionsForMatching(c *gin.Context) (interface{}, error) {
	
	req := requests.GetQuestionsForMatchingRequest{}

	err := c.BindJSON(req)
	if err != nil {
		return nil, err
	}

	// key should come from env
	return h.DiscoverController.GetQuestionsForMatching(req.Offset, os.Getenv("key"))
}
