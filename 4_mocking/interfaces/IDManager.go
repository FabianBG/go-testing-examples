package interfaces

import "github.com/google/uuid"

// mockgen -destination=4_mocking/mocks/mock_IDManager.go
// -package=mocks avalith/testing/4_mocking IDManager

// Interface def

// IDManager Interface definition
type IDManager interface {
	GenID() string
}

// IDManagerLocal local implementation
type IDManagerLocal struct{}

// NewIDManagerLocal constructor
func NewIDManagerLocal() IDManager {
	return &IDManagerLocal{}
}

// GenID implementation
func (o *IDManagerLocal) GenID() string {
	return uuid.New().String()
}
