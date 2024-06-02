gqlent-login:
	PGPASSWORD=postgres psql -h 0.0.0.0 -p 5432 -U postgres gqlent

atlas-diff:
	atlas migrate diff ${migration_name} \
		--config "file://atlas.hcl" \
		--env dev

atlas-apply:
	atlas migrate apply \
		--config "file://atlas.hcl" \
		--env dev

atlas-status:
	atlas migrate apply \
		--config "file://atlas.hcl" \
		--env dev

gql-generate:
	go run github.com/99designs/gqlgen generate