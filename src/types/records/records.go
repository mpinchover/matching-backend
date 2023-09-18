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

type Profile struct {
	gorm.Model
	Gender            string
	GenderPreference  *string
	Age               int64
	MinAgePreference  int64
	MaxAgePreference  int64
	MaxDistanceMeters int64
	ProfileLat        float64
	ProfileLng        float64
}
