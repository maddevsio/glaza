#!/bin/bash
docker pull mysql
docker run --name some-mysql -v /usr/local/opt/docker-volumes/mysql:/var/lib/mysql  -e MYSQL_ROOT_PASSWORD= -d mysql:5.7