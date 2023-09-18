package repo

import (
	"matching/src/mappers"
	"matching/src/types/entities"
	"matching/src/types/records"
)

func (r *Repo) CreateProfile(profile *entities.Profile) error {
	val := mappers.ToProfileRecord(profile)
	return r.DB.Create(val).Error
}

func (r *Repo) GetProfileByUserUUID(uuid string) (*entities.Profile, error) {
	result := &records.Profile{}
	out := r.DB.Where("user_uuid = ?", uuid).Find(result)
	if out.Error != nil {
		return nil, out.Error
	}
	if out.RowsAffected == 0 {
		return nil, nil
	}
	return mappers.ToProfileEntity(result), nil
}
