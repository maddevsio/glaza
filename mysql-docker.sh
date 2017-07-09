#!/bin/bash
#docker pull mysql
docker stop mysql
docker rm mysql
docker run -p 3306:3306 --name mysql -v /usr/local/opt/docker-volumes/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7
docker logs mysql
docker ps