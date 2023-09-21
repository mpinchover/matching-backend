package repo

import (
	"fmt"
	"matching/src/types/entities"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	PAGINATION_MESSAGES              = 20
	PAGINATION_ROOMS                 = 10
	MINIMUM_QUESTION_MATCH_THRESHOLD = 10
)

type IRepo interface {
	CreateTrackedLike(trackedLike *entities.TrackedLike) error
	GetTrackedLikeByUserAndTarget(userUUID, targetUUID string) (*entities.TrackedLike, error)
	UpdateTrackedLike(trackedLike *entities.TrackedLike) error
	UpdateTrackedQuestion(trackedQuestion *entities.TrackedQuestion) error
	CreateTrackedQuestion(trackedQuestion *entities.TrackedQuestion) error
	GetTrackedQuestionByUserAndQuestion(userUUID, targetUUID string) (*entities.TrackedQuestion, error)
	GetCandidateProfilesByMatchedQuestions(questionUUIDs []string, userUUIDsToFilterOut []string, minRequiredMatchThreshold int) ([]*entities.Profile, error)
	GetLikedTrackedQuestionsByUser(userUUID string) ([]*entities.TrackedQuestion, error)
}

type Repo struct {
	DB *gorm.DB
}

func connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"))
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
}

func New() *Repo {
	var db *gorm.DB
	db, err := connect()
	if err != nil {
		panic(err)
	}

	return &Repo{
		DB: db,
	}
}
