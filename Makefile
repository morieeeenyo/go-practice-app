generate-migration:
	docker-compose exec app create destroy -ext sql -dir ./src/database/migrations -seq ${TABLE_NAME}

migrate-up:
	docker-compose exec app migrate -database "mysql://tester:password@tcp(db:3306)/test" -path ./src/database/migrations up

migrate-down:
	docker-compose exec app migrate -database "mysql://tester:password@tcp(db:3306)/test" -path ./src/database/migrations down


# gen-code-from-db:
# 	docker-compose exec app sqlboiler psql --output ./src/database/models --pkgname models --wipe && cd src && go mod tidy
