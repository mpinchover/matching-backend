package discover

import (
	"errors"
	"matching/src/gateways/storage"
	"matching/src/repo"
	"matching/src/types/entities"
	"math"
	"strconv"

	"go.uber.org/fx"
)

type DiscoverController struct {
	repo    repo.IRepo
	storage storage.IStorage
}

type Params struct {
	fx.In

	Repo    *repo.Repo
	Storage *storage.Storage
}

func New(p Params) *DiscoverController {
	return &DiscoverController{
		repo:    p.Repo,
		storage: p.Storage,
	}
}

func (c *DiscoverController) GetQuestionsForMatching(offset int64, key string) ([]*entities.Question, error) {
	df, err := c.storage.GetQuestionsFromStorage(key)
	if err != nil {
		return nil, err
	}

	questions := make([]*entities.Question, len(df.Rows())-1)
	for i, row := range df.Rows() {

		text, ok := row["text"]
		if !ok {
			return nil, errors.New("missing text for question")
		}

		index, ok := row["index"]
		if !ok {
			return nil, errors.New("missing index for question")
		}

		category, ok := row["category"]
		if !ok {
			return nil, errors.New("missing category for question")
		}

		uuid, ok := row["uuid"]
		if !ok {
			return nil, errors.New("missing uuid for question")
		}

		idx, err := strconv.Atoi(*index)
		if err != nil {
			return nil, err
		}
		question := &entities.Question{
			UUID:     *uuid,
			Text:     *text,
			Category: *category,
			Index:    int64(idx),
		}
		questions[i] = question
	}

	if int(offset) >= len(questions) {
		offset = 0
	}

	end := math.Min(float64(offset+20), float64(len(questions)))
	questions = questions[offset:int(end)]
	// get all the tracked questions for these questions that are being chosen
	// so you can see how many were liked. TODO - add it into the zip code later?
	return questions, nil
}
