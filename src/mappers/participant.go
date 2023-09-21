package mappers

import (
	"matching/src/types/entities"
	"matching/src/types/records"
)

func ToParticipantEntities(in []*records.Participant) []*entities.Participant {
	out := make([]*entities.Participant, len(in))

	for i, p := range in {
		out[i] = ToParticipantEntity(p)
	}
	return out
}

func ToParticipantRecords(in []*entities.Participant) []*records.Participant {
	out := make([]*records.Participant, len(in))

	for i, p := range in {
		out[i] = ToParticipantRecord(p)
	}
	return out
}

func ToParticipantEntity(m *records.Participant) *entities.Participant {
	return &entities.Participant{
		UserUUID: m.UserUUID,
		// MatchUUID: m.MatchUUID,
		MatchID: m.MatchID,
	}
}

func ToParticipantRecord(m *entities.Participant) *records.Participant {
	return &records.Participant{
		UserUUID: m.UserUUID,
		// MatchUUID: m.MatchUUID,
		MatchID: m.MatchID,
	}
}
