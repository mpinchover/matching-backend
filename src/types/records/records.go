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
	Liked        bool
}

type TrackedLike struct {
	gorm.Model
	UUID       string
	UserUUID   string
	TargetUUID string
	Liked      bool
}
