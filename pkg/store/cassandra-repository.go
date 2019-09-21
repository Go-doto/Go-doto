package store

import (
	"context"
	dota_api "github.com/Go-doto/Go-doto/pkg/dota-api"
	"github.com/gocql/gocql"
)

type matchRepo struct {
	session *gocql.Session
}

func NewCassandraMatchRepository(session *gocql.Session) MatchRepositoryInterface {
	return matchRepo{}
}

func (m matchRepo) FindByMatchId(ctx context.Context, id dota_api.MatchId) (dota_api.MatchResult, error) {
	return dota_api.MatchResult{}, nil
}

func (m matchRepo) FindByMatchSeqNo(ctx context.Context, id dota_api.MatchSequenceNo) (dota_api.MatchResult, error) {
	return dota_api.MatchResult{}, nil
}

func (m matchRepo) CreateMatch(ctx context.Context, match dota_api.MatchResult) error {
	return nil
}
