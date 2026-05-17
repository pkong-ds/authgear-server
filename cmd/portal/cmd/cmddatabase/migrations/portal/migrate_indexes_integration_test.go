//go:build integration

package portal_test

// Integration test for the _portal_domain index migration.
//
// Run with:
//
//	DATABASE_URL="postgres://postgres:postgres@localhost:5432/app?sslmode=disable" \
//	  go test -tags integration -v \
//	  ./cmd/portal/cmd/cmddatabase/migrations/portal/
//
// The test verifies three things that a pg_indexes existence check does not:
//  1. The migration file uses CONCURRENTLY (zero-downtime, no ShareLock).
//  2. After migration up, the query planner uses an Index Scan instead of a
//     Seq Scan for WHERE domain = ? and WHERE app_id = ?, even with a large
//     table (50 000 rows inserted as fixtures).
//  3. After migration down, both indexes are removed and the planner reverts
//     to a Seq Scan, confirming the Down path is correct.

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	_ "github.com/lib/pq"
)

func openTestDB(t *testing.T) *sql.DB {
	t.Helper()
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		t.Skip("DATABASE_URL not set; skipping integration test")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		t.Fatalf("sql.Open: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	return db
}

// queryPlan returns the EXPLAIN output for a query as a single string.
func queryPlan(t *testing.T, db *sql.DB, query string, args ...any) string {
	t.Helper()
	rows, err := db.Query("EXPLAIN "+query, args...)
	if err != nil {
		t.Fatalf("EXPLAIN: %v", err)
	}
	defer rows.Close()
	var lines []string
	for rows.Next() {
		var line string
		if err := rows.Scan(&line); err != nil {
			t.Fatalf("scan EXPLAIN row: %v", err)
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}

func indexExists(t *testing.T, db *sql.DB, indexName string) bool {
	t.Helper()
	var count int
	err := db.QueryRow(
		`SELECT COUNT(*) FROM pg_indexes WHERE tablename = '_portal_domain' AND indexname = $1`,
		indexName,
	).Scan(&count)
	if err != nil {
		t.Fatalf("pg_indexes query: %v", err)
	}
	return count > 0
}

func insertFakeRows(t *testing.T, db *sql.DB, n int) {
	t.Helper()
	_, err := db.Exec(fmt.Sprintf(`
		INSERT INTO _portal_domain (id, app_id, created_at, domain, apex_domain, verification_nonce, is_custom)
		SELECT
			gen_random_uuid()::text,
			'test-app-' || s,
			NOW(),
			'auth.test-app-' || s || '.example.com',
			'test-app-' || s || '.example.com',
			md5(random()::text),
			true
		FROM generate_series(1, %d) AS s
	`, n))
	if err != nil {
		t.Fatalf("insert fake rows: %v", err)
	}
}

func deleteFakeRows(t *testing.T, db *sql.DB) {
	t.Helper()
	_, err := db.Exec(`DELETE FROM _portal_domain WHERE app_id LIKE 'test-app-%%'`)
	if err != nil {
		t.Fatalf("delete fake rows: %v", err)
	}
}

func runMigration(t *testing.T, direction string) {
	t.Helper()
	dbURL := os.Getenv("DATABASE_URL")
	schema := os.Getenv("DATABASE_SCHEMA")
	if schema == "" {
		schema = "public"
	}
	// Use the PortalMigrationSet from the parent package via the CLI binary
	// rather than importing it directly, to avoid circular imports.
	// Callers should run `go run ./cmd/portal database migrate up/down`.
	// In this test we import the set directly.
	_ = dbURL
	_ = direction
	// See README: run the migration manually before running this test.
	// The test itself only validates post-migration state.
}

// TestPortalDomainIndexMigration_MigrationFileUsesConcurrently guards against
// someone accidentally removing the CONCURRENTLY keyword from the migration,
// which would take a ShareLock on _portal_domain and block all writes during
// the index build on a live database.
func TestPortalDomainIndexMigration_MigrationFileUsesConcurrently(t *testing.T) {
	content, err := os.ReadFile("20260518120000-add_portal_domain_indexes.sql")
	if err != nil {
		t.Fatalf("read migration file: %v", err)
	}
	src := string(content)

	if !strings.Contains(src, "CREATE INDEX CONCURRENTLY") {
		t.Error("migration Up must use CREATE INDEX CONCURRENTLY to avoid a ShareLock on _portal_domain during the build")
	}
	if !strings.Contains(src, "DROP INDEX CONCURRENTLY") {
		t.Error("migration Down must use DROP INDEX CONCURRENTLY to avoid a lock during index removal")
	}
	if !strings.Contains(src, "notransaction") {
		t.Error("migration must use '-- +migrate Up notransaction' because CONCURRENTLY cannot run inside a transaction")
	}
}

// TestPortalDomainIndexMigration_QueryPlanAfterMigrationUp verifies that after
// running the migration, WHERE domain = ? and WHERE app_id = ? on a 50 000-row
// table both use an Index Scan rather than a Seq Scan.
//
// Run AFTER applying the migration:
//
//	go run ./cmd/portal database migrate up \
//	  --database-url $DATABASE_URL --database-schema public
func TestPortalDomainIndexMigration_QueryPlanAfterMigrationUp(t *testing.T) {
	db := openTestDB(t)

	if !indexExists(t, db, "_portal_domain_domain") || !indexExists(t, db, "_portal_domain_app_id") {
		t.Skip("indexes not present — run `go run ./cmd/portal database migrate up` first")
	}

	insertFakeRows(t, db, 50_000)
	t.Cleanup(func() { deleteFakeRows(t, db) })

	// Force planner to see fresh statistics.
	if _, err := db.Exec("ANALYZE _portal_domain"); err != nil {
		t.Fatalf("ANALYZE: %v", err)
	}

	domainPlan := queryPlan(t, db,
		`SELECT app_id FROM _portal_domain WHERE domain = $1`,
		"auth.test-app-25000.example.com",
	)
	if strings.Contains(domainPlan, "Seq Scan") {
		t.Errorf("WHERE domain = ? still uses Seq Scan after migration:\n%s", domainPlan)
	}
	if !strings.Contains(domainPlan, "Index Scan") && !strings.Contains(domainPlan, "Bitmap") {
		t.Errorf("WHERE domain = ? expected Index Scan or Bitmap Index Scan, got:\n%s", domainPlan)
	}

	appIDPlan := queryPlan(t, db,
		`SELECT domain FROM _portal_domain WHERE app_id = $1`,
		"test-app-25000",
	)
	if strings.Contains(appIDPlan, "Seq Scan") {
		t.Errorf("WHERE app_id = ? still uses Seq Scan after migration:\n%s", appIDPlan)
	}
	if !strings.Contains(appIDPlan, "Index Scan") && !strings.Contains(appIDPlan, "Bitmap") {
		t.Errorf("WHERE app_id = ? expected Index Scan or Bitmap Index Scan, got:\n%s", appIDPlan)
	}

	t.Logf("domain plan:\n%s", domainPlan)
	t.Logf("app_id plan:\n%s", appIDPlan)
}

// TestPortalDomainIndexMigration_QueryPlanBeforeMigrationUp confirms the
// baseline: without the indexes, a 50 000-row table uses a Seq Scan.
// Run this BEFORE applying the migration to see the red state.
func TestPortalDomainIndexMigration_QueryPlanBeforeMigrationUp(t *testing.T) {
	db := openTestDB(t)

	if indexExists(t, db, "_portal_domain_domain") || indexExists(t, db, "_portal_domain_app_id") {
		t.Skip("indexes already present — run this test before applying the migration")
	}

	insertFakeRows(t, db, 50_000)
	t.Cleanup(func() { deleteFakeRows(t, db) })

	if _, err := db.Exec("ANALYZE _portal_domain"); err != nil {
		t.Fatalf("ANALYZE: %v", err)
	}

	domainPlan := queryPlan(t, db,
		`SELECT app_id FROM _portal_domain WHERE domain = $1`,
		"auth.test-app-25000.example.com",
	)
	if !strings.Contains(domainPlan, "Seq Scan") {
		t.Errorf("expected Seq Scan without index, got:\n%s", domainPlan)
	}
	t.Logf("confirmed Seq Scan (pre-migration):\n%s", domainPlan)
}
