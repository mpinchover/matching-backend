package requests

import (
	"matching/src/types/entities"
)

type SaveTrackedQuestionRequest struct {
	TrackedQuestion *entities.TrackedQuestion
}

type SaveTrackedQuestionResponse struct {
}

type DeleteMatchRequest struct {
	MatchUUID string
}

type DeleteMatchResponse struct {
}

// deprioritize
type BlockUserRequest struct {
	UserUUID string
}

type BlockUserResponse struct {
}

// shouldn't be an API request
// type CreateMatchRequest struct{}

// type CreateMatchResponse struct{}

type SaveTrackedLikeRequest struct {
	TrackedLike *entities.TrackedLike
}

type SaveTrackedLikeResponse struct{}

type GetQuestionsForMatchingRequest struct {
	Offset int64
}

type GetQuestionsForMatchingResponse struct{}

type APIErrorResponse struct {
	Message string
}
