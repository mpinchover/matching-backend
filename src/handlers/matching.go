package handlers

import (
	"matching/src/controllers/matching"
	"matching/src/types/entities"
	"matching/src/types/requests"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	MatchingController *matching.MatchingController
}

type Params struct {
	fx.In

	MatchingController *matching.MatchingController
}

func New(p Params) *Handler {
	return &Handler{
		MatchingController: p.MatchingController,
	}
}

// delete the room as well
func (h *Handler) DeleteMatch(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// add validation

	return nil, nil
}

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

	req := &requests.CreateTrackedLikeRequest{}

	err := c.BindJSON(req)
	if err != nil {
		return nil, err
	}

	return nil, h.MatchingController.SaveTrackedLike(c, req.TrackedLike)
}

func (m *Handler) UpdateTrackedQuestion(q *entities.TrackedQuestion) error {
	return nil
}
