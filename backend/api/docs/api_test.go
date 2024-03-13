package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/perebaj/policycraft"
	"github.com/perebaj/policycraft/postgres"
)

// MockStorage is a mock implementation of the Storage interface
type MockStorage struct{}

// SavePolicy is a mock implementation of the SavePolicy method
func (m *MockStorage) SavePolicy(_ policycraft.Policy) error {
	return nil
}

func (m *MockStorage) Policies() ([]postgres.Policy, error) {
	return nil, nil
}

// NewMockStorage returns a new instance of MockStorage
func NewMockStorage() *MockStorage {
	return &MockStorage{}
}

func TestSavePolicyHandler(t *testing.T) {
	tests := []struct {
		name     string
		policy   interface{}
		expected int
	}{
		{
			name: "Valid policy",
			policy: policycraft.Policy{
				ID:          uuid.NewString(),
				Name:        "test",
				Value:       1,
				Criteria:    ">=",
				SuccessCase: true,
				Priority:    1,
			},
			expected: http.StatusOK,
		},
		{
			name: "Invalid criteria",
			policy: policycraft.Policy{
				ID:          uuid.NewString(),
				Name:        "test",
				Value:       1,
				Criteria:    "invalid",
				SuccessCase: true,
				Priority:    1,
			},
			expected: http.StatusBadRequest,
		},
		{
			name: "Invalid UUID",
			policy: policycraft.Policy{
				ID:          "invalid",
				Name:        "test",
				Value:       1,
				Criteria:    ">=",
				SuccessCase: true,
				Priority:    1,
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
	handler := SavePolicyHandler(db)

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
				t.Fatalf("expected status code %d, got %d | response: %s", test.expected, w.Code, w.Body.String())
			}
		})
	}
}

func TestPolicyValidateCriteria(t *testing.T) {
	tests := []struct {
		name     string
		criteria string
		wantErr  bool
	}{
		{
			name:     "greater than",
			criteria: ">",
			wantErr:  false,
		},
		{
			name:     "less than",
			criteria: "<",
			wantErr:  false,
		},
		{
			name:     "greater than or equal",
			criteria: ">=",
			wantErr:  false,
		},
		{
			name:     "less than or equal",
			criteria: "<=",
			wantErr:  false,
		},
		{
			name:     "equal",
			criteria: "==",
			wantErr:  false,
		},
		{
			name:     "invalid",
			criteria: "invalid",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Policy{
				Criteria: tt.criteria,
			}
			if err := p.validateCriteria(); (err != nil) != tt.wantErr {
				t.Errorf("Policy.ValidateCriteria() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
