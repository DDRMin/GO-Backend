env "local" {
  src = "file://internal/adapters/migrations/schema.hcl"
  url = getenv("DB_URL")
  dev = "docker://postgres/17/dev?search_path=public"
}
