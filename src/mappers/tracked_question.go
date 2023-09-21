package mappers

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

func ToTrackedQuestionRecords(in []*entities.TrackedQuestion) []*records.TrackedQuestion {
	out := make([]*records.TrackedQuestion, len(in))
	for i, tq := range in {
		out[i] = ToTrackedQuestionRecord(tq)
	}
	return out
}

func ToTrackedQuestionRecord(q *entities.TrackedQuestion) *records.TrackedQuestion {
	return &records.TrackedQuestion{
		UUID:          q.UUID,
		UserUUID:      q.UserUUID,
		QuestionText:  q.QuestionText,
		QuestionIndex: q.QuestionIndex,
		Category:      q.Category,
		QuestionUUID:  q.QuestionUUID,
	}
}
func ToTrackedQuestionEntities(in []*records.TrackedQuestion) []*entities.TrackedQuestion {
	out := make([]*entities.TrackedQuestion, len(in))
	for i, tq := range in {
		out[i] = ToTrackedQuestionEntity(tq)
	}
	return out
}

func ToTrackedQuestionEntity(q *records.TrackedQuestion) *entities.TrackedQuestion {
	return &entities.TrackedQuestion{
		UUID:          q.UUID,
		UserUUID:      q.UserUUID,
		QuestionText:  q.QuestionText,
		QuestionIndex: q.QuestionIndex,
		Category:      q.Category,
		QuestionUUID:  q.QuestionUUID,
	}
}
