FROM php:7.2-cli-stretch

RUN apt-get update && \
    apt-get install libssl-dev zip wget unzip curl nano git -y && \
    rm -rf /var/lib/apt/lists/* && \
    curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN pecl install mongodb \
    && docker-php-ext-enable mongodb

RUN mkdir -p /usr/src/app
RUN wget -O /usr/src/app/phpunit https://phar.phpunit.de/phpunit-7.phar
RUN chmod +x /usr/src/app/phpunit
RUN /usr/src/app/phpunit --version

COPY ./Instagram/composer.json /usr/src/app/composer.json
WORKDIR /usr/src/app
RUN composer install --no-interaction
COPY ./Instagram /usr/src/app
CMD [ "php", "./get.php" ]