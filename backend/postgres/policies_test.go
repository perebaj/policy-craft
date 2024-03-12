//go:build integration
// +build integration

// go: build integration

// using build tag to separate integration tests from unit tests. In that way, we can reduce the time of the test some specific part of the code.
package postgres_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/perebaj/policycraft"
	"github.com/perebaj/policycraft/postgres"
)

// OpenDB create a new database for testing and return a connection to it.
// The idea is to create a new database for each test, so we can run the tests avoiding any kind of conflict between them.
func OpenDB(t *testing.T) *sqlx.DB {
	t.Helper()

	cfg := postgres.Config{
		URL:             os.Getenv("POLICY_CRAFT_POSTGRES_URL"),
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxIdleTime: 1 * time.Minute,
	}

	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		t.Fatalf("error connecting to Postgres: %v", err)
	}
	// The timeout here is a
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if err != nil {
		t.Fatalf("error pinging postgres: %v", err)
	}

	// create a new database with random suffix
	postgresURL, err := url.Parse(cfg.URL)
	if err != nil {
		t.Fatalf("error parsing Postgres connection URL: %v", err)
	}
	database := strings.TrimLeft(postgresURL.Path, "/")

	randSuffix := fmt.Sprintf("%x", time.Now().UnixNano())

	database = fmt.Sprintf("%s-%x", database, randSuffix)
	_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE "%s"`, database))
	if err != nil {
		t.Fatalf("error creating database for test: %v", err)
	}

	postgresURL.Path = "/" + database
	cfg.URL = postgresURL.String()
	testDB, err := postgres.OpenDB(cfg)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// after run the tests, drop the database
	t.Cleanup(func() {
		testDB.Close()
		defer db.Close()
		_, err = db.Exec(fmt.Sprintf(`DROP DATABASE "%s" WITH (FORCE);`, database))
		if err != nil {
			t.Fatalf("error dropping database for test: %v", err)
		}
	})

	return testDB
}

// TestSavePolicy test the SavePolicy method from the Storage struct.
func TestStorageSavePolicy(t *testing.T) {
	db := OpenDB(t)
	defer db.Close()

	storage := postgres.NewStorage(db)
	uuidPolicy := uuid.New()
	policy := policycraft.Policy{
		ID:       uuidPolicy,
		Name:     "policy 1",
		Criteria: ">",
		Value:    1,
	}

	err := storage.SavePolicy(policy)
	if err != nil {
		t.Fatalf("error saving policy: %v", err)
	}

	var got []postgres.Policy
	err = db.Select(&got, "SELECT * FROM policies")
	if err != nil {
		t.Fatalf("error getting count of policies: %v", err)
	}

	if len(got) != 1 {
		t.Fatalf("expected 1 policy, got %d", len(got))
	}

	assert(t, got[0].ID, policy.ID)
	assert(t, got[0].Name, policy.Name)
	assert(t, got[0].Criteria, policy.Criteria)
	assert(t, got[0].Value, policy.Value)

	policy2 := policycraft.Policy{
		ID:       uuidPolicy,
		Name:     "policy 1",
		Criteria: "<",
		Value:    2,
	}

	err = storage.SavePolicy(policy2)
	if err != nil {
		t.Fatalf("error saving policy: %v", err)
	}

	var got2 []postgres.Policy
	err = db.Select(&got2, "SELECT * FROM policies")
	if err != nil {
		t.Fatalf("error getting count of policies: %v", err)
	}

	// Validating if the updated_at field was updated when the same policy was saved, but with different data.
	if len(got) == 1 {
		assert(t, got2[0].ID, policy2.ID)
		assert(t, got2[0].Name, policy2.Name)
		assert(t, got2[0].Criteria, policy2.Criteria)
		assert(t, got2[0].Value, policy2.Value)
		if got2[0].UpdatedAt.Before(got[0].UpdatedAt) {
			t.Errorf("expected updated_at to be after %v, got %v", got2[0].UpdatedAt, got[0].UpdatedAt)
		}
	}
}

func assert(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}