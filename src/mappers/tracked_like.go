package mappers

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

func ToTrackedLikeRecord(tl *entities.TrackedLike) *records.TrackedLike {
	return &records.TrackedLike{
		UUID:       tl.UUID,
		UserUUID:   tl.UserUUID,
		TargetUUID: tl.TargetUUID,
		Liked:      tl.Liked,
	}
}

func ToTrackedLikeEntity(tl *records.TrackedLike) *entities.TrackedLike {
	return &entities.TrackedLike{
		UUID:       tl.UUID,
		UserUUID:   tl.UserUUID,
		TargetUUID: tl.TargetUUID,
		Liked:      tl.Liked,
	}
}
