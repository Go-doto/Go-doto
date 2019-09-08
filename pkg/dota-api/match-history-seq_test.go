package dota_api

import (
	"fmt"
	"github.com/Go-doto/Go-doto/pkg/dota-api/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMatchHistoryBySequenceNum(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, mocks.GetMatchHistoryBySeqNoMock())
	}))
	defer ts.Close()

	apiURL = ts.URL + "/%s/%s"
	client, _ := NewClientWithToken("123")
	details, err := GetMatchHistoryBySequenceNum(client, 0, 0)
	if err == nil {
		t.Error("expected error when matchId is empty")
	}
	details, err = GetMatchHistoryBySequenceNum(client, MatchSequenceNo(41242), 100)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if numberOfResults := len(details.MatchesResult); numberOfResults != 10 {
		t.Errorf("expected 10 result but given %d", numberOfResults)
	}
}
