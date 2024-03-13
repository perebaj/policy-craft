// Package api provides the handlers for the API endpoints
package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/perebaj/policycraft"
	"github.com/perebaj/policycraft/postgres"
)

// Storage is the interface that wraps the postgres methods that iteract with the API
type Storage interface {
	SavePolicy(policy policycraft.Policy) error
	Policies() ([]postgres.Policy, error)
}

type Policy struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Value       *int   `json:"value,omitempty"`
	Criteria    string `json:"criteria"`
	SuccessCase *bool  `json:"success_case,omitempty"`
	Priority    *int   `json:"priority,omitempty"`
}

// validateCriteria checks if the criteria is valid.
func (p *Policy) validateCriteria() error {
	switch p.Criteria {
	case ">", "<", ">=", "<=", "==":
		return nil
	default:
		return fmt.Errorf("invalid criteria: %s", p.Criteria)
	}
}

// Validate checks if the policy is valid. If all required fields are present. Or if their values are equal to the expected.
func (p *Policy) Validate() error {
	var errs []error

	//valida if id is a valid UUID
	if _, err := uuid.Parse(p.ID); err != nil {
		errs = append(errs, fmt.Errorf("id is not a valid UUID"))
	}
	if p.Value == nil {
		errs = append(errs, fmt.Errorf("value is required"))
	}
	if p.SuccessCase == nil {
		errs = append(errs, fmt.Errorf("success_case is required"))
	}
	if p.Priority == nil {
		errs = append(errs, fmt.Errorf("priority is required"))
	}
	if p.validateCriteria() != nil {
		errs = append(errs, fmt.Errorf("invalid criteria"))
	}
	return errors.Join(errs...)
}

// SavePolicyHandler returns a http.HandlerFunc that receive a policy and save it to the database
func SavePolicyHandler(db Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var policy Policy
		err := json.NewDecoder(r.Body).Decode(&policy)
		if err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}
		err = policy.Validate()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		p := policycraft.Policy{
			ID:          policy.ID,
			Name:        policy.Name,
			Value:       *policy.Value,
			Criteria:    policy.Criteria,
			SuccessCase: *policy.SuccessCase,
			Priority:    *policy.Priority,
		}

		err = db.SavePolicy(p)
		if err != nil {
			http.Error(w, "failed to save policy", http.StatusInternalServerError)
			return
		}
	}
}

// ListPoliciesHandler returns a http.HandlerFunc that get all the policies from the database and return it as a response
func ListPoliciesHandler(db Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		policies, err := db.Policies()
		if err != nil {
			slog.Error("failed to get policies", "error", err)
			http.Error(w, "failed to get policies", http.StatusInternalServerError)
			return
		}

		//convert the policies to response body
		policiesByte, err := json.Marshal(policies)
		if err != nil {
			http.Error(w, "failed to marshal policies", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(policiesByte)
	}
}
