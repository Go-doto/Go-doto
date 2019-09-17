package dota_api

import "fmt"

type ValidationError struct {
	Reason string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error: %s", e.Reason)
}

// can retry
type ServerError struct {
	Reason string
	Code   int
}

func (e ServerError) Error() string {
	return fmt.Sprintf("server error: %s with code %d", e.Reason, e.Code)
}

type AccessForbiddenError struct {
	Reason string
}

func (e AccessForbiddenError) Error() string {
	return fmt.Sprintf("access forbidden: %s", e.Reason)
}

type UnknownError struct {
	Reason string
	Inner  error
}

func (e UnknownError) Error() string {
	if e.Inner != nil {
		return fmt.Sprintf("unknown error: %s: %v", e.Reason, e.Inner)
	}
	return fmt.Sprintf("unknown error: %s", e.Reason)
}
func (e UnknownError) Unwrap() error {
	return e.Inner
}

// can retry
type HttpClientError struct {
	Inner error
}

func (e HttpClientError) Error() string {
	return fmt.Sprintf("http client error: %v", e.Inner)
}

func (e HttpClientError) Unwrap() error {
	return e.Inner
}
