#!/bin/sh

docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD="1234" --name postgres postgres:alpine
sleep 2

docker run \
 --rm \
 -d \
 -p 8080:8080 \
 -p 8088:8088  \
 -e ODJ_DEP_ADDRESS_BOOK_API_URI=http://localhost:8080  \
 -e ODJ_DEP_POSTGRESQL_URL=postgresql://postgres:1234@host.docker.internal:5432/postgres \
 -e ODJ_DEP_USER_NAME='Hans Hubert' \
 --name addressbook \
 hoehne/addressbook:2.0

open "http://localhost:8088"
while :
do
  sleep 10 &
  wait $!
done

docker kill addressbook
docker kill postgres
