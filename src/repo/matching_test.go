package repo

import (
	"fmt"
	"matching/src/types/entities"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type RepoSuite struct {
	suite.Suite

	repo   *Repo
	mainDB *gorm.DB
}

func (r *RepoSuite) SetupSuite() {

	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_HOST"), "3306", os.Getenv("MYSQL_DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		// FullSaveAssociations: true,
	})
	r.NoError(err)
	r.repo = &Repo{}
	r.mainDB = db

}

func (r *RepoSuite) SetupTest() {
	r.repo.DB = r.mainDB.Begin()

}

func (r *RepoSuite) TeardownSuite() {
	// r.repo.DB.Rollback()
}

func TestRepoSuite(t *testing.T) {
	suite.Run(t, new(RepoSuite))
}

func (s *RepoSuite) TestGetCandidateProfiles() {
	defer s.repo.DB.Rollback()

	user := &entities.Profile{
		Gender:           "MAN",
		GenderPreference: "FEMALE",
		Age:              24,
		MinAgePreference: 30,
		MaxAgePreference: 40,
		ProfileLat:       20,
		ProfileLng:       30,
		UserUUID:         uuid.New().String(),
	}

	err := s.repo.CreateProfile(user)
	s.NoError(err)

	q1UUID := uuid.New().String()
	q2UUID := uuid.New().String()
	q3UUID := uuid.New().String()
	q4UUID := uuid.New().String()
	q5UUID := uuid.New().String()
	q6UUID := uuid.New().String()
	q7UUID := uuid.New().String()
	q8UUID := uuid.New().String()

	trackedQuestions := []*entities.TrackedQuestion{
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      user.UserUUID,
			QuestionUUID:  q1UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      user.UserUUID,
			QuestionUUID:  q2UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      user.UserUUID,
			QuestionUUID:  q3UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      user.UserUUID,
			QuestionUUID:  q4UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      user.UserUUID,
			QuestionUUID:  q5UUID,
		},
	}

	for _, tq := range trackedQuestions {
		err := s.repo.CreateTrackedQuestion(tq)
		s.NoError(err)
	}

	// candidate1 - should match with 90% match
	candidate1 := &entities.Profile{
		Gender:           "MAN",
		GenderPreference: "FEMALE",
		Age:              24,
		MinAgePreference: 30,
		MaxAgePreference: 40,
		ProfileLat:       20,
		ProfileLng:       30,
		UserUUID:         uuid.New().String(),
	}

	err = s.repo.CreateProfile(candidate1)
	s.NoError(err)

	trackedQuestions = []*entities.TrackedQuestion{
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate1.UserUUID,
			QuestionUUID:  q2UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate1.UserUUID,
			QuestionUUID:  q3UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate1.UserUUID,
			QuestionUUID:  q4UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate1.UserUUID,
			QuestionUUID:  q5UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate1.UserUUID,
			QuestionUUID:  q6UUID,
		},
	}

	for _, tq := range trackedQuestions {
		err := s.repo.CreateTrackedQuestion(tq)
		s.NoError(err)
	}

	// candidate2 - should match with 90% match
	candidate2 := &entities.Profile{
		Gender:           "MAN",
		GenderPreference: "FEMALE",
		Age:              24,
		MinAgePreference: 30,
		MaxAgePreference: 40,
		ProfileLat:       20,
		ProfileLng:       30,
		UserUUID:         uuid.New().String(),
	}

	err = s.repo.CreateProfile(candidate2)
	s.NoError(err)

	trackedQuestions = []*entities.TrackedQuestion{
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate2.UserUUID,
			QuestionUUID:  q4UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate2.UserUUID,
			QuestionUUID:  q5UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate2.UserUUID,
			QuestionUUID:  q6UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate2.UserUUID,
			QuestionUUID:  q7UUID,
		},
		{
			UUID:          uuid.New().String(),
			QuestionText:  "QuestionText",
			QuestionIndex: 0,
			Category:      "CATEGORY",
			UserUUID:      candidate2.UserUUID,
			QuestionUUID:  q8UUID,
		},
	}

	for _, tq := range trackedQuestions {
		err := s.repo.CreateTrackedQuestion(tq)
		s.NoError(err)
	}

	userAnsweredQuestions := []string{q1UUID, q2UUID, q3UUID, q4UUID, q5UUID}
	profiles, err := s.repo.GetCandidateProfiles(userAnsweredQuestions, []string{user.UserUUID}, 3)
	s.NoError(err)
	s.NotNil(profiles)
	s.Len(profiles, 1)

	userAnsweredQuestions = []string{q1UUID, q2UUID, q3UUID, q4UUID, q5UUID}
	profiles, err = s.repo.GetCandidateProfiles(userAnsweredQuestions, []string{user.UserUUID}, 2)
	s.NoError(err)
	s.NotNil(profiles)
	s.Len(profiles, 2)

	userAnsweredQuestions = []string{q1UUID, q2UUID, q3UUID, q4UUID, q5UUID}
	profiles, err = s.repo.GetCandidateProfiles(userAnsweredQuestions, []string{user.UserUUID}, 4)
	s.NoError(err)
	s.NotNil(profiles)
	s.Len(profiles, 1)

	userAnsweredQuestions = []string{q1UUID, q2UUID, q3UUID, q4UUID, q5UUID}
	profiles, err = s.repo.GetCandidateProfiles(userAnsweredQuestions, []string{user.UserUUID}, 5)
	s.NoError(err)
	s.Nil(profiles)
	s.Len(profiles, 0)

	userAnsweredQuestions = []string{q1UUID, q2UUID, q3UUID, q4UUID, q5UUID}
	profiles, err = s.repo.GetCandidateProfiles(userAnsweredQuestions, nil, 0)
	s.NoError(err)
	s.NotNil(profiles)
	s.Len(profiles, 3)

	// should still return the user profile
	userAnsweredQuestions = []string{q1UUID, q2UUID, q3UUID, q4UUID, q5UUID}
	profiles, err = s.repo.GetCandidateProfiles(userAnsweredQuestions, nil, 5)
	s.NoError(err)
	s.NotNil(profiles)
	s.Len(profiles, 1)
}

func (s *RepoSuite) TestCreateTrackedQuestion() {
	defer s.repo.DB.Rollback()

	trackedQuestion := &entities.TrackedQuestion{
		UUID:          uuid.New().String(),
		QuestionText:  "question text",
		QuestionIndex: 10,
		QuestionUUID:  uuid.New().String(),
		UserUUID:      uuid.New().String(),
		Category:      "CATEGORY",
	}

	err := s.repo.CreateTrackedQuestion(trackedQuestion)
	s.NoError(err)

	tq, err := s.repo.GetTrackedQuestionByUserAndQuestion(trackedQuestion.UserUUID, trackedQuestion.QuestionUUID)
	s.NoError(err)
	s.NotNil(tq)
	s.Equal(trackedQuestion, tq)

	tq, err = s.repo.GetTrackedQuestionByUserAndQuestion("other", trackedQuestion.QuestionUUID)
	s.NoError(err)
	s.Nil(tq)
}
