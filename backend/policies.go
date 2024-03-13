// Package policycraft have the core logic for the service policycraft. Including the entry point entities that arrive from the client.
// policies.go gather the entiy policy and some operations related to it.
package policycraft

// Policy is the struct that represent the business entity policy. In other words, it is the struct that will be used as input of the service.
type Policy struct {
	// ID is the unique identifier of the policy.
	ID string `json:"id" db:"id"`
	// Name is the name of the policy.
	Name string `json:"name" db:"name"`
	// Criteria is the criteria that will be used to compare the value. It can be: >, <, >=, <=, ==.
	Criteria string `json:"criteria" db:"criteria"`
	// Value is the value that will be used to compare with the criteria.
	Value int `json:"value" db:"value"`
	// SuccessCase is the boolean that will be used to compare the result of the policy
	SuccessCase bool `json:"success_case" db:"success_case"`
	// Priority is the priority of the policy. The lower the number, the higher the priority.
	Priority int `json:"priority" db:"priority"`
}
