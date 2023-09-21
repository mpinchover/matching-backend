package records

import (
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type TrackedQuestion struct {
	gorm.Model
	UUID          string
	QuestionText  string
	QuestionIndex int64
	Category      string
	UserUUID      string
	QuestionUUID  string
}

// // TableName overrides the table name used by SeenBy to `seen_by`
// func (TrackedQuestion) TableName() string {
// 	return "tracked_questions"
// }

type TrackedLike struct {
	gorm.Model
	UUID       string
	UserUUID   string
	TargetUUID string
	Liked      bool
}

// maybe just use rooms in messaging server
type Match struct {
	gorm.Model
	UUID     string
	RoomUUID string
}

type Profile struct {
	gorm.Model
	Gender            string
	GenderPreference  string
	Age               int64
	MinAgePreference  int64
	MaxAgePreference  int64
	MaxDistanceMeters int64
	ProfileLat        float64
	ProfileLng        float64
	UserUUID          string `gorm:"user_uuid"`
}
