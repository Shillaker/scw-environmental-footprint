# HTTPS and SSL certs

- Deploy to the VM
- SSH into the machine
- Run `export NGINX_CONF_DIR=./nginx/conf-ssl/`
- Run `docker-compose down && docker-compose up -d`

This will set up NGINX in a mode to respond to challenges without HTTPS redirect.

Then run the following:

```
docker-compose run --rm  certbot certonly --webroot --webroot-path /var/www/certbot/ -d scw-impact.simonshillaker.com
```

Then

- Run `unset NGINX_CONF_DIR`
- Run `docker-compose down && docker-compose up -d`

Check the result at https://scw-impact.simonshillaker.com

## Links

- [Blog](https://mindsers.blog/post/https-using-nginx-certbot-docker/) on using docker-compose and certbot
