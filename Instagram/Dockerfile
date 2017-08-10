FROM php:7.0-cli
RUN docker-php-ext-install mysqli
RUN apt-get update && \
    apt-get install curl nano git -y && \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
COPY ./Instagram /usr/src/app
COPY ./lib /usr/src/lib
WORKDIR /usr/src/app
RUN composer install --no-interaction
CMD [ "php", "./get.php" ]
