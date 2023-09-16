package handlers

import (
	"matching/src/controllers/matching"
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

func (h *Handler) CreateMatch(c *gin.Context) (interface{}, error) {
	// add validation

	return nil, nil
}

func (h *Handler) CreateTrackedLike(c *gin.Context) (interface{}, error) {
	// add validation

	req := &requests.CreateTrackedLikeRequest{}

	err := c.BindJSON(req)
	if err != nil {
		return nil, err
	}

	return nil, h.MatchingController.CreateTrackedLike(c, req)
}

func (m *Handler) UpdateTrackedQuestion(q *requests.TrackedQuestion) error {
	return nil
}

// func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) (interface{}, error) {
// 	// TODO - validate the create room request
// 	req := &requests.CreateRoomRequest{}
// 	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
// 		return nil, err
// 	}

// 	ctx := r.Context()

// 	err := validation.ValidateRequest(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	room, err := h.ControlTowerCtrlr.CreateRoom(ctx, req.Members)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &requests.CreateRoomResponse{
// 		Room: room,
// 	}, nil
// }
