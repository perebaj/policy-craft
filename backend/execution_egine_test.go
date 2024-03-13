package policycraft

import "testing"

func TestEvaluate(t *testing.T) {
	// Create a new execution
	execution := &Execution{
		CustomFields: map[string]interface{}{
			"age": 16,
		},
	}

	// Create a new policy
	policy := Policy{
		Name:        "age",
		Criteria:    ">",
		Value:       17,
		SuccessCase: false,
		Priority:    1,
	}

	// Evaluate the policy
	result, err := execution.Evaluate([]Policy{policy})
	if err != nil {
		t.Errorf("Error evaluating policy: %s", err)
	}

	if result != true {
		t.Errorf("The policy should be evaluated as true")
	}

	// Create a new policy
	policy2 := Policy{
		Name:        "rank",
		Criteria:    ">",
		Value:       15,
		SuccessCase: false,
		Priority:    2,
	}

	execution.CustomFields["age"] = 18
	execution.CustomFields["rank"] = 16

	// Evaluate the policy
	result, err = execution.Evaluate([]Policy{policy, policy2})
	if err != nil {
		t.Errorf("Error evaluating policy: %s", err)
	}

	if result != false {
		t.Errorf("The policy should be evaluated as false")
	}

	// creating a new policy

	policy3 := Policy{
		Name:        "income",
		Criteria:    "==",
		Value:       1000,
		SuccessCase: true,
		Priority:    3,
	}

	execution.CustomFields["income"] = 1000

	// Evaluate the policy
	result, err = execution.Evaluate([]Policy{policy, policy2, policy3})
	if err != nil {
		t.Errorf("Error evaluating policy: %s", err)
	}

	if result != true {
		t.Errorf("The policy should be evaluated as true")
	}

	// creating a new policy where the custom field is not present

	policy4 := Policy{
		Name:        "size",
		Criteria:    "==",
		Value:       1,
		SuccessCase: true,
		Priority:    4,
	}

	// Evaluate the policy
	_, err = execution.Evaluate([]Policy{policy, policy2, policy3, policy4})
	if err == nil {
		t.Errorf("Expecting an error evaluating policy because the custom field is not present")
	}

	// Expecting an error when sending an empty policy list
	_, err = execution.Evaluate([]Policy{})
	if err == nil {
		t.Errorf("Expecting an error evaluating policy because the policy list is empty")
	}

	// Expecting an error when I send a customField key that is not present in the policies
	execution.CustomFields["notpresent"] = 1
	_, err = execution.Evaluate([]Policy{policy, policy2, policy3})

	if err == nil {
		t.Errorf("Expecting an error evaluating policy because the custom field is not present in the policies")
	}
}
