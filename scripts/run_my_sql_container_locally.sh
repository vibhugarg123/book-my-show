#!/usr/bin/env bash
result=$(docker ps -f "name=full_db_mysql" | wc -l)
if [ $result -gt 1 ]; then
  docker stop full_db_mysql
  docker rm full_db_mysql
fi

docker run --name full_db_mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ons_vg -d mysql
##CREATE DATABASE BOOK_MY_SHOW
mysql -u root -h 127.0.0.1 -pons_vg -e "CREATE DATABASE IF NOT EXISTS BOOK_MY_SHOW;"
