package parser

import (
	"github.com/stretchr/testify/mock"

	service "github.com/geoirb/face-search/internal/face-search"
)

// Mock ...
type Mock struct {
	mock.Mock
}

// GetProfileList ...
func (m *Mock) GetProfileList(payload []byte) ([]service.Profile, error) {
	args := m.Called(payload)
	if p, ok := args.Get(0).([]service.Profile); ok {
		return p, args.Error(1)
	}
	return nil, nil
}
