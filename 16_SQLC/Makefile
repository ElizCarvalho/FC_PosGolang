create-migration:
	migrate create -ext=sql -dir=sql/migrations -seq init

migrate-up: 
	migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" up

migrate-down:
	migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/courses" down

validate-migration:
	docker compose exec mysql mysql -uroot -p courses -e "SHOW TABLES;"

.PHONY: migrate-up migrate-down create-migration validate-migration