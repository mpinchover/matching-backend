package repo

import (
	"matching/src/mappers"
	"matching/src/types/entities"
	"matching/src/types/records"
)

func (r *Repo) CreateTrackedLike(trackedLike *entities.TrackedLike) error {
	return r.DB.Save(trackedLike).Error
}

// TODO - use the maps for more security
func (r *Repo) UpdateTrackedLike(trackedLike *entities.TrackedLike) error {

	// updates := map[string]interface{}{
	// 	"user_uuid":   trackedLike.UserUUID,
	// 	"target_uuid": trackedLike.TargetUUID,
	// 	"liked":       trackedLike.Liked,
	// }
	return r.DB.Where("uuid = ?", trackedLike.UUID).Updates(trackedLike).Error
}

func (r *Repo) GetTrackedLikeByUserAndTarget(userUUID, targetUUID string) (*entities.TrackedLike, error) {
	trackedLike := &records.TrackedLike{}
	err := r.DB.Where("user_uuid = ?", userUUID).Where("target_uuid = ?", targetUUID).Find(&trackedLike).Error
	if err != nil {
		return nil, err
	}
	return mappers.ToTrackedLikeEntity(trackedLike), nil
}

func (r *Repo) UpdateTrackedQuestion(trackedQuestion *entities.TrackedQuestion) error {
	tq := mappers.ToTrackedQuestionRecord(trackedQuestion)
	err := r.DB.Where("user_uuid = ?", tq.UserUUID).
		Where("question_uuid = ?", tq.QuestionUUID).
		Update("messages", tq).Error
	return err
}

func (r *Repo) CreateTrackedQuestion(trackedQuestion *entities.TrackedQuestion) error {
	tq := mappers.ToTrackedQuestionRecord(trackedQuestion)
	return r.DB.Save(tq).Error
}

func (r *Repo) GetTrackedQuestionByUserAndQuestion(userUUID, targetUUID string) (*entities.TrackedQuestion, error) {
	tq := &records.TrackedQuestion{}
	err := r.DB.Where("user_uuid = ?", userUUID).Where("target_uuid = ?", targetUUID).Find(&tq).Error
	if err != nil {
		return nil, err
	}
	return mappers.ToTrackedQuestionEntity(tq), nil
}
