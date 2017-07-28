FROM php:7.0-cli
RUN docker-php-ext-install mysqli
COPY ./Instagram /usr/src/app
COPY ./lib /usr/src/lib
WORKDIR /usr/src/app
CMD [ "php", "./get.php" ]
