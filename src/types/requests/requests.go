package requests

import (
	"matching/src/types/entities"
)

type UpdateTrackedQuestionRequest struct {
	TrackedQuestion *entities.TrackedQuestion
}

type UpdateTrackedQuestionResponse struct {
}

type CreateTrackedQuestionRequest struct {
	TrackedQuestion *entities.TrackedQuestion
}

type CreateTrackedQuestionResponse struct {
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

type CreateTrackedLikeRequest struct{}

type CreateTrackedLikeResponse struct{}

type GetQuestionsForMatchingRequest struct{}

type GetQuestionsForMatchingResponse struct{}

type APIErrorResponse struct {
	Message string
}
