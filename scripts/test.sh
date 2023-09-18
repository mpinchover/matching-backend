#!/bin/bash
mysql -uroot -proot --protocol=tcp -h localhost --port=3310 -e "
    drop database if exists matching_test_db;
    create database matching_test_db;
"
# redis-cli flushall
redis-cli -h localhost -p 6380 flushall

# running this from root dir
mysql -uroot -proot --protocol=tcp -h localhost --port=3310 matching_test_db < ./database/schema.sql
docker exec -it matching-backend-matchingserver-1 bash -c "cd src && go test ./... -v -count=1 -failfast"
