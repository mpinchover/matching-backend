package repo

import (
	"matching/src/types/entities"

	"github.com/google/uuid"
)

func (s *RepoSuite) TestGetCandidateProfiles() {
	defer s.repo.DB.Rollback()

	// create a profile
	profile := &entities.Profile{
		Gender:            "MAN",
		GenderPreference:  "WOMAN",
		Age:               25,
		MinAgePreference:  30,
		MaxAgePreference:  40,
		UserUUID:          uuid.New().String(),
		MaxDistanceMeters: 10,
		ProfileLat:        100,
		ProfileLng:        120,
	}

	err := s.repo.CreateProfile(profile)
	s.NoError(err)

	existingProfile, err := s.repo.GetProfileByUserUUID(profile.UserUUID)
	s.NoError(err)
	s.NotNil(existingProfile)
	s.Equal(profile, existingProfile)

	existingProfile, err = s.repo.GetProfileByUserUUID("something-else")
	s.NoError(err)
	s.Nil(existingProfile)
}
