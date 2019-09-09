package dota_api

import (
	"errors"
)

func GetMatchDetails(client ClientInterface, matchId MatchId) (MatchResult, error) {
	matchDetails := MatchResult{}
	if matchId == 0 {
		return matchDetails, errors.New("required matchId")
	}

	resp, err := client.MakeRequest("GetMatchDetails", map[string]string{
		"match_id": matchId.ToString(),
	})
	if err != nil {
		return matchDetails, err
	}
	err = matchDetails.UnmarshalJSON(resp.Result)

	if err != nil {
		return matchDetails, err
	}

	if matchDetails.Error != "" {
		return matchDetails, errors.New(matchDetails.Error)
	}
	if matchDetails.MatchId == 0 {
		return matchDetails, errors.New("match not parsed")
	}
	return matchDetails, err
}
