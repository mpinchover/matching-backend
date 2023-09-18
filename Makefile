run-api:
	docker compose up

stop:
	docker compose down

migrate:
	mysql -uroot -proot --protocol=tcp -h localhost --port=3310 matching < ./db/schema.sql

migrate-test:
	mysql -uroot -proot --protocol=tcp -h localhost --port=3310 matching_test_db < ./db/schema.sql