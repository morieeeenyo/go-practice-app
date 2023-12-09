generate-migration:
	migrate create -ext sql -dir ./backend/src/database/migrations -seq ${TABLE_NAME}

migrate-up:
	migrate -database "mysql://root@tcp(127.0.0.1:3306)/go_practice_app_development" -path ./backend/src/database/migrations up

migrate-down:
	migrate -database "mysql://root@tcp(127.0.0.1:3306)/go_practice_app_development" -path ./backend/src/database/migrations down

gen-code-from-db:
	sqlboiler mysql --output ./backend/src/database/models --pkgname models --wipe --config ./backend/config.toml && cd src && go mod tidy

start-frontend:
	cd frontend && yarn dev

start-backend:
	cd backend && air -c .air.toml