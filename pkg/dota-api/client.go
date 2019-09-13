package dota_api

/*
client for Dota 2 api.

	client, _ := api.NewClientWithToken("token_string")
	resp, err := api.GetMatchDetails(client, "4949341670")
	if err != nil {
		log.Fatal(err)
		os.Exit(0)
	}
	fmt.Println(resp)
*/
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	apiURLProduction = "https://api.steampowered.com/IDOTA2Match_570/%s/%s"
	apiVersion       = "V001"
	requestTimeout   = 10
)

// default url
var apiURL = apiURLProduction

type ClientInterface interface {
	MakeRequest(method string, params map[string]string) (APIResponse, error)
}

// client struct
type client struct {
	Token  string
	Client *http.Client
}

// Create new client instance from string token
// Default client has 10 seconds timeout.
func NewClientWithToken(token string) (*client, error) {
	if token == "" {
		return &client{}, errors.New("token is required")
	}
	httpClient := http.Client{
		Timeout: requestTimeout * time.Second,
	}
	return &client{
		Token:  token,
		Client: &httpClient,
	}, nil
}

// Create Make request to specific steam API method with get query params.
func (c *client) MakeRequest(method string, params map[string]string) (APIResponse, error) {
	// create request url
	endpoint := fmt.Sprintf(apiURL, method, apiVersion)
	// create request instance than add query params.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Print(err)
		return APIResponse{}, err
	}

	// build query params
	q := req.URL.Query()
	q.Add("key", c.Token)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.Client.Do(req)
	// send request
	defer resp.Body.Close()
	if err != nil {
		return APIResponse{}, err
	}

	if resp.StatusCode > 400 {
		return APIResponse{}, errors.New(fmt.Sprintf("Server error. Code %d", resp.StatusCode))
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return APIResponse{}, err
	}

	var apiResponse APIResponse
	json.Unmarshal(body, &apiResponse)
	return apiResponse, nil
}
