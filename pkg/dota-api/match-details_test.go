package dota_api

import (
	"fmt"
	"github.com/Go-doto/Go-doto/pkg/dota-api/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMatchDetails(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, mocks.GetMatchDetailsMock())
	}))
	defer ts.Close()

	apiURL = ts.URL + "/%s/%s"
	client, _ := NewClientWithToken("123")
	details, err := GetMatchDetails(client, "")
	if err == nil {
		t.Error("expected error when matchId is empty")
	}
	details, err = GetMatchDetails(client, "41242")
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if details.MatchId != 4949341670 {
		t.Errorf("invalid parsed match id. expected %s given %d", "4294967295", details.MatchId)
	}
}
