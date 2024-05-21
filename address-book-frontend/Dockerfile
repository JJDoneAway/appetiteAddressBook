FROM nginx:alpine

RUN apk upgrade --no-cache

COPY nginx.conf /etc/nginx/nginx.conf
COPY env_injector /usr/share/env_injector
COPY .env.odj /usr/share/.env.odj
COPY dist/ /usr/share/nginx/html
COPY .well-known/ /usr/share/nginx/html/.well-known/
COPY /node_modules/oidc-client/lib/oidc-client.min.js /usr/share/nginx/html/node_modules/oidc-client/lib/

RUN chmod 777 /usr/share/env_injector

CMD ["/bin/sh",  "-c",  "/usr/share/env_injector /usr/share/nginx/html/ /usr/share/.env.odj && exec nginx -g 'daemon off;'"]
