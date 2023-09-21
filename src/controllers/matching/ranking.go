package matching

import (
	"matching/src/types/entities"
	"matching/src/types/enums"
)

const (
	MIN_REQUIRED_QUESTIONS_ANSWERED = 30
)

func (m *MatchingController) FindCompatibleCandidates(trackedQuestion *entities.TrackedQuestion) (*entities.MatchesForUserResult, error) {
	result := &entities.MatchesForUserResult{}

	// get questions this user has liked
	// todo - validate a user uuid exists
	likedQuestionsByUser, err := m.repo.GetLikedTrackedQuestionsByUser(trackedQuestion.UserUUID)
	if err != nil {
		return nil, err
	}
	if len(likedQuestionsByUser) < MIN_REQUIRED_QUESTIONS_ANSWERED {
		result.AbortCode = enums.ABORT_CODE_BELOW_MIN_REQUIRED_QUESTIONS_ANSWERED.String()
		return result, nil
	}

	uuids := make([]string, len(likedQuestionsByUser))
	for i, tq := range likedQuestionsByUser {
		uuids[i] = tq.UUID
	}

	// TODO - make a filter so you can also filter by matching prefs as well
	userUUIDSToFilterOut := []string{}
	candidates, err := m.repo.GetCandidateProfilesByMatchedQuestions(uuids, userUUIDSToFilterOut, 30)
	if err != nil {
		return nil, err
	}
	if len(candidates) == 0 {
		result.AbortCode = enums.ABORT_CODE_NO_MATCHING_CANDIDATES.String()
		return result, nil
	}
	result.Candidates = candidates
	return result, nil
}
