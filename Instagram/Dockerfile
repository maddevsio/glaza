FROM php:7.0-cli
RUN docker-php-ext-install mysqli
RUN apt-get update && \
    apt-get install curl nano git -y && \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
COPY ./Instagram/composer.json /usr/src/app/composer.json
WORKDIR /usr/src/app
RUN composer install --no-interaction
COPY ./Instagram /usr/src/app
CMD [ "php", "./get.php" ]