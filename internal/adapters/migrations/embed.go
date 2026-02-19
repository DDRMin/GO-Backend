package migrations

import (
	_ "embed"
	"strings"
)

//go:embed schema.sql
var schema string

// Schema returns the embedded SQL with CREATE TABLE converted to
// CREATE TABLE IF NOT EXISTS, making it safe to run on every startup
// without conflicting with Atlas-generated schema.sql.
func Schema() string {
	return strings.ReplaceAll(schema, "CREATE TABLE ", "CREATE TABLE IF NOT EXISTS ")
}
