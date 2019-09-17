package dota_api

func GetMatchDetails(client ClientInterface, matchId MatchId) (MatchResult, error) {
	matchDetails := MatchResult{}
	if matchId == 0 {
		return matchDetails, ValidationError{"required matchId"}
	}

	resp, err := client.MakeRequest("GetMatchDetails", map[string]string{
		"match_id": matchId.ToString(),
	})
	if err != nil {
		return matchDetails, err
	}
	err = matchDetails.UnmarshalJSON(resp.Result)

	if err != nil {
		return matchDetails, UnknownError{Inner: err, Reason: "unmarshal error"}
	}

	if matchDetails.Error != "" {
		return matchDetails, ValidationError{matchDetails.Error}
	}
	if matchDetails.MatchId == 0 {
		return matchDetails, ValidationError{"match not parsed. Validate data"}
	}
	return matchDetails, err
}
