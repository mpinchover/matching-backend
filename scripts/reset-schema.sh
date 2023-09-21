#!/bin/bash
mysql -uroot -proot --protocol=tcp -h localhost --port=3310 -e "
    drop database if exists matching_test_db;
    create database matching_test_db;
"
# running this from root dir
mysql -uroot -proot --protocol=tcp -h localhost --port=3310 matching_test_db < ./database/schema.sql