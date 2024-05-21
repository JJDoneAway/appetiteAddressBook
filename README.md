# Showcase for Appetite
This is a very basic address book application consisting out of
* Frontend in vue js
* backend in go
* Postgres as DB

It is deployed within one single Docker container which is very unusual but for friar enough for the POC

# how to build
run `build_decker.sh`

# how to run 
run `run_docker.sh`

# used environment parameters

```
ODJ_DEP_ADDRESS_BOOK_API_URI=http://{your public IP}:8080
ODJ_DEP_POSTGRESQL_URL=postgresql://{db user}:{db password}@{db host name}:{db port}/{db name}
ODJ_DEP_USER_NAME='Name of the address book owner'
```

