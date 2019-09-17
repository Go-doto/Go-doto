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
	Token       string
	Client      *http.Client
	rateLimiter *rateLimiter
}

type rateLimiter struct {
	requestsCount   int
	lastRequestTime time.Time
}

// Create new client instance from string token
// Default client has 10 seconds timeout.
func NewClientWithToken(token string) (*client, error) {
	if token == "" {
		return &client{}, ValidationError{"token is required"}
	}
	httpClient := http.Client{
		Timeout: requestTimeout * time.Second,
	}
	return &client{
		Token:       token,
		Client:      &httpClient,
		rateLimiter: &rateLimiter{},
	}, nil
}

// Create Make request to specific steam API method with get query params.
func (c *client) MakeRequest(method string, params map[string]string) (APIResponse, error) {
	c.rateLimiter.wait()
	// create request url
	endpoint := fmt.Sprintf(apiURL, method, apiVersion)
	// create request instance than add query params.
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Print(err)
		return APIResponse{}, HttpClientError{err}
	}

	// build query params
	q := req.URL.Query()
	q.Add("key", c.Token)
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := c.Client.Do(req)
	if err != nil {
		return APIResponse{}, HttpClientError{err}
	}
	// send request
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		if resp.StatusCode == 403 {
			return APIResponse{}, AccessForbiddenError{Reason: "Access is denied. Retrying will not help. Please verify your token"}
		} else {
			return APIResponse{}, ServerError{Reason: "Server error", Code: resp.StatusCode}
		}
	}
	c.rateLimiter.update()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return APIResponse{}, UnknownError{Inner: err, Reason: "read from response body error"}
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return APIResponse{}, UnknownError{Inner: err, Reason: "json unmarshal error"}
	}
	return apiResponse, nil
}

// sleep function for rate limiter
func (s *rateLimiter) wait() {
	if s.requestsCount == 5 {
		secs := time.Since(s.lastRequestTime).Seconds()
		ms := int((1 - secs) * 1000)
		if ms > 0 {
			duration := time.Duration(ms * int(time.Millisecond))
			log.Println("Alarm! 3 requests per second. Sleeping for", ms, "ms")
			time.Sleep(duration)
		}

		s.requestsCount = 0
	}
}

// every request updates rateLimiter struct
func (s *rateLimiter) update() {
	s.requestsCount++
	s.lastRequestTime = time.Now()
}
