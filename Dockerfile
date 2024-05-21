FROM nginx:alpine

RUN apk upgrade --no-cache

#backend
RUN date > timestamp.txt
COPY address-book-backend/app /app
COPY address-book-backend/.env /.env


# fronend
COPY address-book-frontend/nginx.conf /etc/nginx/nginx.conf
COPY address-book-frontend/env_injector /usr/share/env_injector
COPY address-book-frontend/.env.odj /usr/share/.env.odj
COPY address-book-frontend/dist/ /usr/share/nginx/html
COPY address-book-frontend/.well-known/ /usr/share/nginx/html/.well-known/

RUN chmod 777 /usr/share/env_injector

#start script
COPY run_address_book.sh /run_address_book.sh
RUN chmod 777 /run_address_book.sh


CMD ["/run_address_book.sh"]
