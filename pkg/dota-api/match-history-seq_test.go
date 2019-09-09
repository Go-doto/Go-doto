package dota_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var jsonData []byte

func init() {
	// Open our jsonFile
	jsonFile, err := os.Open("mocks/full_json_seq_no.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonData, _ = ioutil.ReadAll(jsonFile)
}

func TestGetMatchHistoryBySequenceNum(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(jsonData))
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
	if numberOfResults := len(details.MatchesResult); numberOfResults != 100 {
		t.Errorf("expected 10 result but given %d", numberOfResults)
	}
}

//without easyjson
func BenchmarkGetMatchHistoryBySequenceNum(b *testing.B) {
	response := APIResponse{}
	_ = json.Unmarshal(jsonData, &response)
	for i := 0; i < b.N; i++ {
		matchHistory := MatchHistoryBySequenceNo{}
		_ = json.Unmarshal(response.Result, &matchHistory)
	}
}

//with easyjson
func BenchmarkGetMatchHistoryBySequenceNumEasyJson(b *testing.B) {
	response := APIResponse{}
	_ = json.Unmarshal(jsonData, &response)
	for i := 0; i < b.N; i++ {
		matchHistory := MatchHistoryBySequenceNo{}
		_ = matchHistory.UnmarshalJSON(response.Result)
	}
}
