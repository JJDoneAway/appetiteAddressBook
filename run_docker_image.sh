#!/bin/sh

docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD="1234" --name postgres postgres:alpine
sleep 2
docker run --rm -p 8080:8080 -p 8088:8088  -e ODJ_DEP_ADDRESS_BOOK_API_URI=http://localhost:8080  -e ODJ_DEP_POSTGRESQL_URL=postgresql://postgres:1234@host.docker.internal:5432/postgres --name addressbook addressbook 

docker kill addressbook
docker kill postgres
