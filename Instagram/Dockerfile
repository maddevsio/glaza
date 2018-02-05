FROM php:7.2-cli-stretch

RUN apt-get update && \
    apt-get install libssl-dev curl nano git -y && \ 
    rm -rf /var/lib/apt/lists/* && \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN pecl install mongodb \
    && docker-php-ext-enable mongodb

COPY ./Instagram/composer.json /usr/src/app/composer.json
WORKDIR /usr/src/app
RUN composer install --no-interaction
COPY ./Instagram /usr/src/app
CMD [ "php", "./get.php" ]