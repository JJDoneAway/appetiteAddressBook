# Showcase for Appetite
This is a very basic address book application consisting out of
* Frontend in vue js
* backend in go
* Postges as DB

It is deployed within one single Docker container which is very unusual but for friar enough for the POC

# 

docker build -t addressbook .    

docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD="1234" --name postgres postgres:alpine
docker run --rm -p 8080:8080 -p 8088:8088  -e ODJ_DEP_ADDRESS_BOOK_API_URI=http://localhost:8080  -e ODJ_DEP_POSTGRESQL_URL=postgresql://postgres:1234@host.docker.internal:5432/postgres --name addressbook addressbook   


