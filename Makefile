run-api:
	docker compose up

stop:
	docker compose down

migrate:
	mysql -uroot -proot --protocol=tcp -h localhost --port=3310 matching < ./database/schema.sql

migrate-test:
	mysql -uroot -proot --protocol=tcp -h localhost --port=3310 matching_test_db < ./database/schema.sql

test:
	./scripts/test.sh