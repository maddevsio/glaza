how to generate a password for basic auth

https://www.digitalocean.com/community/tutorials/how-to-set-up-password-authentication-with-nginx-on-ubuntu-14-04

You can add a username to the file using this command. We are using sammy as our username, but you can use whatever name you'd like:
```
sudo sh -c "echo -n 'sammy:' >> /etc/nginx/.htpasswd"
```
Next, add an encrypted password entry for the username by typing:
```
sudo sh -c "openssl passwd -apr1 >> /etc/nginx/.htpasswd"
```