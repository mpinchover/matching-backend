package matching

import "go.uber.org/fx"

type MatchingController struct {
}

type Params struct {
	fx.In
}

func New(p Params) *MatchingController {
	return &MatchingController{}
}
