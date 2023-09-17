package mappers

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

func ToTrackedQuestionRecord(q *entities.TrackedQuestion) *records.TrackedQuestion {
	return &records.TrackedQuestion{
		UUID:         q.UUID,
		UserUUID:     q.UserUUID,
		Text:         q.Text,
		Index:        q.Index,
		Category:     q.Category,
		QuestionUUID: q.QuestionUUID,
	}
}

func ToTrackedQuestionEntity(q *records.TrackedQuestion) *entities.TrackedQuestion {
	return &entities.TrackedQuestion{
		UUID:         q.UUID,
		UserUUID:     q.UserUUID,
		Text:         q.Text,
		Index:        q.Index,
		Category:     q.Category,
		QuestionUUID: q.QuestionUUID,
	}
}
