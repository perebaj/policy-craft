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
)

// Storage is the interface that wraps the postgres methods that iteract with the API
type Storage interface {
	SavePolicy(policy policycraft.Policy) error
	Policies() ([]policycraft.Policy, error)
}

// Policy is the struct that represents the policy entity in the API.
type Policy struct {
	// ID is the unique identifier for the policy.
	ID string `json:"id"`
	// Name is the name of the policy.
	Name string `json:"name"`
	// Value is the value that the policy will use to compare.
	Value *int `json:"value,omitempty"`
	// Criteria is the criteria that the policy will use to compare the value.
	Criteria string `json:"criteria"`
	// SuccessCase is the boolean that will be used to compare the result of the policy
	SuccessCase *bool `json:"success_case,omitempty"`
	// Priority is the priority of the policy. The lower the number, the higher the priority.
	Priority *int `json:"priority,omitempty"`
	// IMPORTANT: The pointer fields were chosen to be able to differentiate between the absence of the field and the zero value of the field.
}

// validateCriteria checks if the criteria field is valid.
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = policy.Validate()
		if err != nil {
			sendErr(w, err.Error(), http.StatusBadRequest)
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
			slog.Error("failed to save policy", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}

// ListPoliciesHandler returns a http.HandlerFunc that get all the policies from the database and return it as a response
func ListPoliciesHandler(db Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		policies, err := db.Policies()
		if err != nil {
			slog.Error("failed to get policies", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//convert the policies to response body
		policiesByte, err := json.Marshal(policies)
		if err != nil {
			slog.Error("failed to marshal policies", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		defer func() {
			_, _ = w.Write(policiesByte)
		}()
	}
}

// ExecutionEngineHandler returns a http.HandlerFunc that receive a custom fields and evaluate the policies
func ExecutionEngineHandler(db Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var e policycraft.Execution
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		policies, err := db.Policies()
		if err != nil {
			slog.Error("failed to get policies", "error", err)
			sendErr(w, "failed to get policies", http.StatusInternalServerError)
			return
		}

		decision, err := e.Evaluate(policies)
		if err != nil {
			slog.Error("failed to evaluate policies", "error", err)
			sendErr(w, "failed to evaluate policies "+err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		defer func() {
			_, _ = w.Write([]byte(fmt.Sprintf(`{"decision": %t}`, decision)))
		}()
	}
}

// ErrMsg is the struct that represents the error message in the API.
type ErrMsg struct {
	Msg string `json:"msg"`
}

// sendErr send a error message as a response
func sendErr(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errMsg := ErrMsg{Msg: msg}
	errMsgBytest, err := json.Marshal(errMsg)
	if err != nil {
		http.Error(w, "failed to marshal error message", http.StatusInternalServerError)
		return
	}
	w.Write(errMsgBytest)
}
