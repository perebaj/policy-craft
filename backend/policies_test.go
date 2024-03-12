package policycraft

import "testing"

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
			if err := p.ValidateCriteria(); (err != nil) != tt.wantErr {
				t.Errorf("Policy.ValidateCriteria() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
