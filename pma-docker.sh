#!/usr/bin/env bash
docker stop pma
docker rm pma
docker run --name pma -d --link mysql:db -p 8080:80 phpmyadmin/phpmyadmin
docker ps -a
sleep 2
docker logs pma
