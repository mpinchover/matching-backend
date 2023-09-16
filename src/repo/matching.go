package repo

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

// first get the tracked like to see if its there
func (r *Repo) CreateTrackedLike(trackedLike *entities.TrackedLike) error {
	return r.DB.Save(trackedLike).Error
}

func (r *Repo) GetTrackedLikeByUserAndTarget(userUUID, targetUUID string) (entities.TrackedLike, error) {
	trackedLike := &records.TrackedLike{}
	err := r.DB.Where("user_uuid = ?", userUUID).Where("target_uuid = ?", targetUUID).Find(&trackedLike).Error
	return tackedLike, err
}

func (r *Repo) UpdateTrackedQuestion(trackedQuestion *entities.TrackedLike) {
	err := r.DB.Where("user_uuid = ?", trackedQuestion.UserUUID).
		Where("question_uuid = ?", trackedQuestion.QuestionUUID).
		Update("messages", trackedQuestion).Error
}
