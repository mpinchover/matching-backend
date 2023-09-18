package mappers

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

type Profile struct {
	Gender            string
	GenderPreference  *string
	Age               int64
	MinAgePreference  int64
	MaxAgePreference  int64
	MaxDistanceMeters int64
	ProfileLat        float64
	ProfileLng        float64
}

func ToProfileEntities(profiles []*records.Profile) []*entities.Profile {
	out := make([]*entities.Profile, len(profiles))
	for i, p := range profiles {
		out[i] = ToProfileEntity(p)
	}
	return out
}

func ToProfileEntity(pr *records.Profile) *entities.Profile {
	return &entities.Profile{
		Gender:            pr.Gender,
		GenderPreference:  pr.GenderPreference,
		Age:               pr.Age,
		MinAgePreference:  pr.MinAgePreference,
		MaxAgePreference:  pr.MaxAgePreference,
		MaxDistanceMeters: pr.MaxDistanceMeters,
		ProfileLat:        pr.ProfileLat,
		ProfileLng:        pr.ProfileLng,
	}
}

func ToProfileRecords(profiles []*entities.Profile) []*records.Profile {
	out := make([]*records.Profile, len(profiles))
	for i, p := range profiles {
		out[i] = ToProfileRecord(p)
	}
	return out
}

func ToProfileRecord(pr *entities.Profile) *records.Profile {
	return &records.Profile{
		Gender:            pr.Gender,
		GenderPreference:  pr.GenderPreference,
		Age:               pr.Age,
		MinAgePreference:  pr.MinAgePreference,
		MaxAgePreference:  pr.MaxAgePreference,
		MaxDistanceMeters: pr.MaxDistanceMeters,
		ProfileLat:        pr.ProfileLat,
		ProfileLng:        pr.ProfileLng,
	}
}
