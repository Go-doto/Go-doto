package dota_api

import (
	"errors"
	"strconv"
)

func GetMatchHistoryBySequenceNum(client ClientInterface, fromMatchSeqNo MatchSequenceNo, matchesRequested int) (MatchHistoryBySequenceNo, error) {
	historyData := MatchHistoryBySequenceNo{}
	if fromMatchSeqNo == 0 {
		return historyData, errors.New("required MatchSequenceNo")
	}

	if matchesRequested == 0 {
		matchesRequested = 100
	}

	if matchesRequested < 0 {
		return historyData, errors.New("matchesRequested must be greater than zero")
	}

	resp, err := client.MakeRequest("GetMatchHistoryBySequenceNum", map[string]string{
		"start_at_match_seq_num": fromMatchSeqNo.ToString(),
		"matches_requested":      strconv.Itoa(matchesRequested),
	})
	if err != nil {
		return historyData, err
	}

	err = historyData.UnmarshalJSON(resp.Result)

	if err != nil {
		return historyData, err
	}

	if historyData.Status != 1 {
		return historyData, errors.New(historyData.StatusDetail)
	}

	return historyData, err
}
