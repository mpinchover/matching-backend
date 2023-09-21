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
