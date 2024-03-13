// Package postgres ...
// policies.go gather all the database operations related to the policies entity
package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/perebaj/policycraft"
)

// Storage is the struct that will hold the database connection.
type Storage struct {
	// using sqlx, because this library provides some useful feature on top of the standard library.
	// Documentation: https://jmoiron.github.io/sqlx/
	db *sqlx.DB
}

// NewStorage is the constructor for the Storage struct.
func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{db: db}
}

// Policy is the struct that represents the policy entity in the database.
type Policy struct {
	// ID is the unique identifier for the policy.
	ID uuid.UUID `json:"id" db:"id"`
	// Name is the name of the policy.
	Name string `json:"name" db:"name"`
	// Criteria is the criteria that the policy will use to compare the value.
	Criteria string `json:"criteria" db:"criteria"`
	// Value is the value that the policy will use to compare.
	Value int `json:"value" db:"value"`
	// SuccessCase is the boolean that will be used to compare the result of the policy
	SuccessCase bool `json:"success_case" db:"success_case"`
	// Priority is the priority of the policy. The lower the number, the higher the priority.
	Priority int `json:"priority" db:"priority"`
	// UpdatedAt is the time when the policy was updated.
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// SavePolicy save a policy in the database. If the policy already exists, it will be updated.
func (s *Storage) SavePolicy(policy policycraft.Policy) error {
	_, err := s.db.NamedExec(`
		INSERT INTO policies (id, name, criteria, value, success_case, priority) VALUES (:id, :name, :criteria, :value, :success_case, :priority)
		ON CONFLICT (id) DO UPDATE SET name = :name, criteria = :criteria, value = :value
	`, policy)

	return err
}

// Policies returns all the policies in the database.
func (s *Storage) Policies() ([]Policy, error) {
	var policies []Policy
	err := s.db.Select(&policies, "SELECT * FROM policies ORDER BY priority ASC")
	return policies, err
}
