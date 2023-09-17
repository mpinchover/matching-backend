package records

import (
	"gorm.io/gorm"
)

type TrackedQuestion struct {
	gorm.Model
	UUID         string
	Text         string
	Index        int64
	Category     string
	UserUUID     string
	QuestionUUID string
}

type TrackedLike struct {
	gorm.Model
	UUID       string
	UserUUID   string
	TargetUUID string
	Liked      bool
}

type Match struct {
	gorm.Model
	UUID     string
	RoomUUID string
	// Members  []*Member
}

// type Member struct {
// 	gorm.Model
// 	UserUUID  string
// 	MatchUUID string
// }
