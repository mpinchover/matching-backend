package mappers

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

func ToMatchEntity(m *records.Match) *entities.Match {
	return &entities.Match{
		UUID:         m.UUID,
		RoomUUID:     m.RoomUUID,
		Participants: ToParticipantEntities(m.Participants),
	}
}

func ToMatchRecord(m *entities.Match) *records.Match {
	return &records.Match{
		UUID:         m.UUID,
		RoomUUID:     m.RoomUUID,
		Participants: ToParticipantRecords(m.Participants),
	}
}
