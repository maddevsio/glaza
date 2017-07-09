#!/bin/bash
docker pull mysql
docker run -p 3306:3306 --name mysql -v /usr/local/opt/docker-volumes/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD= -d mysql:5.7