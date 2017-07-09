#!/bin/bash
docker pull mysql
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD= -d mysql:5.7