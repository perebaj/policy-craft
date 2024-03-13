// Package policycraft have the core logic for the service policycraft. Including the entry point entities that arrive from the client.
// policies.go gather the entiy policy and some operations related to it.
package policycraft

import (
	"fmt"
)

// Policy is the struct that will hold the policy data that it's being received from the client.
type Policy struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Criteria string `json:"criteria"`
	Value    int    `json:"value"`
}

// ValidateCriteria checks if the criteria is valid.
func (p *Policy) ValidateCriteria() error {
	switch p.Criteria {
	case ">", "<", ">=", "<=", "==":
		return nil
	default:
		return fmt.Errorf("invalid criteria: %s", p.Criteria)
	}
}
