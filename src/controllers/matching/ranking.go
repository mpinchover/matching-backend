package matching

// func (m *MatchingController) RankCandidates(userUUID string, candidateUUID string) (*matchingTypes.Profile, error) {
// 	// get commonly like questions between user and candidate

// 	commonLikedQuestions, err := m.GetCommonLikedQuestions(userUUID, candidateUUID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// TODO - get better at determining this. Maybe just need really focused questions?
// 	if len(commonLikedQuestions) < 10 {
// 		return nil, nil
// 	}

// 	// show the candidate because it's a potential match

// 	// return the candidate
// 	return &matchingTypes.Profile{
// 		UUID: candidateUUID,
// 	}, nil
// }

// func (m *MatchingController) GetCommonLikedQuestions(userUUID string, candidateUUID string) ([]*matchingTypes.Question, error) {
// 	/*
// 		SELECT column1, column2, ...
// 	FROM your_table
// 	WHERE item IN ('item1', 'item2', 'item3', 'item4', ...)  -- Replace with your array of items
// 	GROUP BY column1, column2, ...
// 	HAVING COUNT(item) >= 3;

// 	*/
// 	userTrackedQuestions, err := m.Repo.GetLikedTrackedQuestionsByUserUUID(userUUID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	uTrackedQuestions := make([]*matchingTypes.TrackedQuestion, len(userTrackedQuestions))
// 	for i, q := range userTrackedQuestions {
// 		reqQuestion := &matchingTypes.TrackedQuestion{
// 			Text:         q.Text,
// 			UUID:         q.UUID,
// 			Category:     q.Category,
// 			Index:        q.Index,
// 			Liked:        q.Liked,
// 			QuestionUUID: q.QuestionUUID,
// 			UserUUID:     q.UserUUID,
// 		}
// 		uTrackedQuestions[i] = reqQuestion
// 	}

// 	// TODO - convert to requests
// 	candidateTrackedQuestions, err := m.Repo.GetLikedTrackedQuestionsByUserUUID(candidateUUID)
// 	cTrackedQuestions := make([]*matchingTypes.TrackedQuestion, len(candidateTrackedQuestions))
// 	for i, q := range candidateTrackedQuestions {
// 		reqQuestion := &matchingTypes.TrackedQuestion{
// 			Text:         q.Text,
// 			UUID:         q.UUID,
// 			Category:     q.Category,
// 			Index:        q.Index,
// 			Liked:        q.Liked,
// 			QuestionUUID: q.QuestionUUID,
// 			UserUUID:     q.UserUUID,
// 		}
// 		cTrackedQuestions[i] = reqQuestion
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	candidateQuestionsMap := map[string]*matchingTypes.TrackedQuestion{}
// 	for _, q := range cTrackedQuestions {
// 		candidateQuestionsMap[q.UUID] = q
// 	}

// 	commonLikes := []*matchingTypes.Question{}
// 	for _, q := range userTrackedQuestions {
// 		if _, ok := candidateQuestionsMap[q.UUID]; !ok {
// 			continue
// 		}

// 		// start if with just things people like
// 		if q.Liked && q.Liked == candidateQuestionsMap[q.UUID].Liked {
// 			question := &matchingTypes.Question{
// 				UUID:     q.QuestionUUID,
// 				Index:    q.Index,
// 				Text:     q.Text,
// 				Category: q.Category,
// 			}
// 			commonLikes = append(commonLikes, question)
// 		}
// 	}

// 	return commonLikes, nil
// }
