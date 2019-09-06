package dota_api

import (
	"encoding/json"
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
	json.Unmarshal(resp.Result, &matchDetails)
	if matchDetails.Error.ErrorMsg != "" {
		return matchDetails, errors.New(matchDetails.Error.ErrorMsg)
	}
	if matchDetails.MatchId == 0 {
		return matchDetails, errors.New("match not parsed")
	}
	return matchDetails, err
}
