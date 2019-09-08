package dota_api

import (
	"encoding/json"
	"errors"
	"strconv"
)

func GetMatchHistoryBySequenceNum(client ClientInterface, matchSeqNo MatchSequenceNo, matchesRequested int) (MatchHistoryBySequenceNo, error) {
	historyData := MatchHistoryBySequenceNo{}
	if matchSeqNo == 0 {
		return historyData, errors.New("required MatchSequenceNo")
	}

	if matchesRequested == 0 {
		matchesRequested = 100
	}

	if matchesRequested < 0 {
		return historyData, errors.New("matchesRequested must be greater than zero")
	}

	resp, err := client.MakeRequest("GetMatchHistoryBySequenceNum", map[string]string{
		"start_at_match_seq_num": matchSeqNo.ToString(),
		"matches_requested":      strconv.Itoa(matchesRequested),
	})
	if err != nil {
		return historyData, err
	}

	json.Unmarshal(resp.Result, &historyData)

	if historyData.Status != 1 {
		return historyData, errors.New(historyData.StatusDetail)
	}

	return historyData, err
}
