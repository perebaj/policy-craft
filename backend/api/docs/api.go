package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/perebaj/policycraft"
)

// Storage is the interface that wraps the postgres methods that iteract with the API
type Storage interface {
	SavePolicy(policy policycraft.Policy) error
}

func CreatePolicyHandler(db Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var policy policycraft.Policy
		err := json.NewDecoder(r.Body).Decode(&policy)
		if err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		err = policy.ValidateCriteria()
		if err != nil {
			http.Error(w, "invalid criteria", http.StatusBadRequest)
			return
		}

		_, err = uuid.Parse(policy.ID)
		if err != nil {
			http.Error(w, "invalid uuid", http.StatusBadRequest)
			return
		}

		err = db.SavePolicy(policy)
		if err != nil {
			http.Error(w, "failed to save policy", http.StatusInternalServerError)
			return
		}
	}
}
