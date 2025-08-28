start_db:
	docker compose up -d

stop_db:
	docker compose down

dbinit:
	docker exec -i gifma-db psql -U gifma_user -d gifma_db < database/init.sql


createdb:
	docker exec -i gifma-db psql -U gifma_user -d gifma_db -c "CREATE DATABASE gifma_db;"


test:
	go test -v -cover ./...


server:
	APP_ENV=dev go run main.go

.PHONY: start_db stop_db dbinit createdb test server