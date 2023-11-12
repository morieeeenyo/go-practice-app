generate-migration:
	docker-compose exec app migrate create -ext sql -dir ./backend/src/database/migrations -seq ${TABLE_NAME}

migrate-up:
	docker-compose exec app migrate -database "mysql://tester:password@tcp(db:3306)/test" -path ./backend/src/database/migrations up

migrate-down:
	docker-compose exec app migrate -database "mysql://tester:password@tcp(db:3306)/test" -path ./backend/src/database/migrations down

gen-code-from-db:
	docker-compose exec app sqlboiler mysql --output ./backend/src/database/models --pkgname models --wipe --config ./backend/config.toml && cd src && go mod tidy

start-frontend:
	cd frontend && yarn dev