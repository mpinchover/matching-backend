package repo

import (
	"matching/src/mappers"
	"matching/src/types/entities"
	"matching/src/types/records"
)

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
