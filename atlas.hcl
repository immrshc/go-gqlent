env "dev" {
  url = "postgres://postgres:postgres@0.0.0.0:5432/gqlent?search_path=public&sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"
  src = "ent://ent/schema"

  migration {
    dir = "file://ent/migrate/migrations"
  }
}
