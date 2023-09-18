package repo

import (
	"fmt"
	"os"
	"testing"

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

// this function executes before the test suite begins execution
func (r *RepoSuite) SetupSuite() {

	dsn := fmt.Sprintf("root:root@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_HOST"), "3306", os.Getenv("MYSQL_DB_NAME"))
	// make connection here
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// FullSaveAssociations: true,
	})
	r.NoError(err)
	r.repo = &Repo{
		// DB: db,
	}
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

// func (r *RepoSuite) TestGetCandidateProfiles() {
// 	defer r.repo.DB.Rollback()

// }
