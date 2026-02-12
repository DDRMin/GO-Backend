env "local" {
  src = "file://internal/adapters/migrations/schema.hcl"
  url = "postgres://user:password@localhost:5432/mydb?sslmode=disable"
  dev = "docker://postgres/17/dev?search_path=public"
}
