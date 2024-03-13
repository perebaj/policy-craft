// Package policycraft ...
// execution_engine.go gather the logic for evaluating policies with given data.
package policycraft

import "fmt"

// Execution is a struct that will be used to evaluate the policies comparing them with an input data(custom fields).
type Execution struct {
	// CustomFields is a map of custom fields that will be used to evaluate the policies
	CustomFields map[string]interface{}
	//Observation: The CustomFields map is dynamic map field that can store any type of data.
	// The duty of this code is to verify if the custom fields that were passed, could be used to evaluate the policies.
}

// Evaluate will evaluate the policies and custom fields and return a boolean value indicating execution decision.
func (e *Execution) Evaluate(policies []Policy) (bool, error) {
	// Isn't possible to evaluate a policy without any policies
	if len(policies) == 0 {
		return false, fmt.Errorf("no policies to evaluate")
	}

	// Validating if all custom fields keys have a respective policy to be evaluated
	policyMap := make(map[string]bool)
	for _, policy := range policies {
		_, ok := e.CustomFields[policy.Name]
		if !ok {
			return false, fmt.Errorf("value '%s' not found in custom fields", policy.Name)
		}
		// loading the policy name into a map to increse the performance of the next validation
		policyMap[policy.Name] = true
	}

	// Validating if there is a custom field that doesn't exist in the policies
	for key := range e.CustomFields {
		_, ok := policyMap[key]
		if !ok {
			return false, fmt.Errorf("the value '%s' doesn't exist in the policies", key)
		}
	}

	// Observations:
	// 1. It's expected that the custom field value is always an int, so we can safely cast it to int
	// 2. We are assuming the policies are ordered by the priority, so we can iterate over them safely
	for _, policy := range policies {
		switch policy.Criteria {
		case ">":
			if !(e.CustomFields[policy.Name].(int) > policy.Value) {
				return !policy.SuccessCase, nil
			}
		case "<":
			if !(e.CustomFields[policy.Name].(int) < policy.Value) {
				return !policy.SuccessCase, nil
			}
		case ">=":
			if !(e.CustomFields[policy.Name].(int) >= policy.Value) {
				return !policy.SuccessCase, nil
			}
		case "<=":
			if !(e.CustomFields[policy.Name].(int) <= policy.Value) {
				return !policy.SuccessCase, nil
			}
		case "==":
			if !(e.CustomFields[policy.Name].(int) == policy.Value) {
				return !policy.SuccessCase, nil
			}
		}
	}
	// If all policies are evaluated as true, we return the last policy success case
	return policies[len(policies)-1].SuccessCase, nil
}
