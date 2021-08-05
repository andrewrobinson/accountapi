package mocks

import (
	"errors"
	"net/http"
)

// MockClient is the mock client
type MockClientThatErrorsOnDo struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Do is the mock client's `Do` func
func (m *MockClientThatErrorsOnDo) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("MOCK ERROR ON DO")
}
