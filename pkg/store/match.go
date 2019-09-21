package store

import (
	"context"
	dota_api "github.com/Go-doto/Go-doto/pkg/dota-api"
)

type MatchRepositoryInterface interface {
	FindByMatchId(ctx context.Context, id dota_api.MatchId) (dota_api.MatchResult, error)
	FindByMatchSeqNo(ctx context.Context, id dota_api.MatchSequenceNo) (dota_api.MatchResult, error)
	CreateMatch(ctx context.Context, match dota_api.MatchResult) error
}
