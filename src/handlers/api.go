package handlers

import (
	"matching/src/controllers/matching"
	"net/http"

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

func (h *Handler) CreateMatch(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// add validation

	return nil, nil
}

func (h *Handler) CreateTrackedLike(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	// add validation

	return nil, nil
}

func (m *Handler) UpdateTrackedQuestion(q *requests.TrackedQuestion) error {

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
