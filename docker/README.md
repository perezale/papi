
## SCANOSS LOAD BALANCER EXAMPLE

Example configuration of NGINX as a Reverse Proxy for the scanning service and SSL management with Certbot.

## Nginx

docker-compose up -d nginx 

## Certbot

docker-compose run --rm  certbot certonly --webroot --webroot-path /var/www/certbot/ --dry-run -d api.scanoss.app

## Miscellaneous

$ export PATH=$PATH:$(go env GOPATH)/bin

## Reference links

https://www.nginx.com/blog/nginx-1-13-10-grpc/

https://caddyserver.com/docs/caddyfile/directives/reverse_proxy

https://caddyserver.com/docs/caddyfile/directives/file_server