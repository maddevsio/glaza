FROM php:7.0-cli
COPY ./Instagram /usr/src/app
COPY ./lib /usr/src/lib
WORKDIR /usr/src/app
RUN [ "php", "./get.php" ]
CMD [ "php", "./get.php" ]
