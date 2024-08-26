migrate:
	@migrate create -ext sql -dir db/migrate/migrations -seq ${filter-out $@,$(MAKECMDGOALS)}

