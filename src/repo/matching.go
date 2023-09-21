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

// userUUIDsToFilterOut - original user uuid, blocked user uuids, currently matched user uuids, etc.
// check first to make sure that theyhave even answered enough questions
func (r *Repo) GetCandidateProfilesByMatchedQuestions(questionUUIDs []string, userUUIDsToFilterOut []string, minRequiredMatchThreshold int) ([]*entities.Profile, error) {
	uuids := []string{}
	params := []interface{}{}
	params = append(params, questionUUIDs)

	query := `
		select user_uuid from tracked_questions
		where question_uuid in (?)
	`

	if len(userUUIDsToFilterOut) > 0 {
		query += ` and user_uuid not in (?) `
		params = append(params, userUUIDsToFilterOut)
	}

	params = append(params, minRequiredMatchThreshold)
	query += ` 
		and deleted_at is null 
		group by user_uuid
		having count(question_uuid) >= ?
	`

	out := r.DB.Raw(query, params...).Scan(&uuids)
	if out.Error != nil {
		return nil, out.Error
	}

	if out.RowsAffected == 0 {
		return nil, nil
	}

	profiles := []*records.Profile{}
	out = r.DB.Where("user_uuid in ?", uuids).Find(&profiles)
	if out.Error != nil {
		return nil, out.Error
	}
	if out.RowsAffected == 0 {
		return nil, nil
	}

	return mappers.ToProfileEntities(profiles), nil
}
