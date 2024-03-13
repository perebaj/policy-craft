package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/perebaj/policycraft"
	api "github.com/perebaj/policycraft/api/docs"
)

// MockStorage is a mock implementation of the Storage interface
type MockStorage struct{}

// SavePolicy is a mock implementation of the SavePolicy method
func (m *MockStorage) SavePolicy(policy policycraft.Policy) error {
	return nil
}

// NewMockStorage returns a new instance of MockStorage
func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func TestCreatePolicyHandler(t *testing.T) {
	tests := []struct {
		name     string
		policy   interface{}
		expected int
	}{
		{
			name: "Valid policy",
			policy: policycraft.Policy{
				ID:       uuid.NewString(),
				Name:     "test",
				Value:    1,
				Criteria: ">=",
			},
			expected: http.StatusOK,
		},
		{
			name: "Invalid criteria",
			policy: policycraft.Policy{
				ID:       uuid.NewString(),
				Name:     "test",
				Value:    1,
				Criteria: "invalid",
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Invalid UUID",
			policy: policycraft.Policy{
				ID:       "invalid",
				Name:     "test",
				Value:    1,
				Criteria: ">=",
			},
			expected: http.StatusBadRequest,
		},
		{
			name:     "Invalid request body",
			policy:   "invalid",
			expected: http.StatusBadRequest,
		},
	}

	db := NewMockStorage()
	handler := api.CreatePolicyHandler(db)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(test.policy)
			if err != nil {
				t.Fatalf("failed to encode policy: %v", err)
			}

			req := httptest.NewRequest("POST", "/policies", &buf)
			w := httptest.NewRecorder()

			handler(w, req)

			if w.Code != test.expected {
				t.Fatalf("expected status code %d, got %d", test.expected, w.Code)
			}
		})
	}
}
