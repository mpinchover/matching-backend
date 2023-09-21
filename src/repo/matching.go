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
	out := r.DB.Where("user_uuid = ?", userUUID).Where("target_uuid = ?", targetUUID).Find(&trackedLike)
	if out.Error != nil {
		return nil, out.Error
	}

	if out.RowsAffected == 0 {
		return nil, nil
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

func (r *Repo) GetTrackedQuestionByUserAndQuestion(userUUID, questionUUID string) (*entities.TrackedQuestion, error) {
	tq := &records.TrackedQuestion{}
	out := r.DB.Where("user_uuid = ?", userUUID).Where("question_uuid = ?", questionUUID).Find(&tq)
	if out.Error != nil {
		return nil, out.Error
	}
	if out.RowsAffected == 0 {
		return nil, nil
	}
	return mappers.ToTrackedQuestionEntity(tq), nil
}

func (r *Repo) GetLikedTrackedQuestionsByUser(userUUID string) ([]*entities.TrackedQuestion, error) {
	tqs := []*records.TrackedQuestion{}
	out := r.DB.Where("user_uuid = ?", userUUID).Find(&tqs)
	if out.Error != nil {
		return nil, out.Error
	}
	if out.RowsAffected == 0 {
		return nil, nil
	}
	return mappers.ToTrackedQuestionEntities(tqs), nil
}

func (r *Repo) CreateMatch(match *entities.Match) error {
	_match := mappers.ToMatchRecord(match)
	return r.DB.Create(_match).Error
}
