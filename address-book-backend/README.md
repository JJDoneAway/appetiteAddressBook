# Fork of https://dev.azure.com/schwarzit-wiking/schwarzit.johannes-meets-stackit/_git/johannes-meets-stackit.monorepo?path=/address-book-backend

!!! Mind. This is just a demo and not mend to be deployed in any scenario eave close to production. !!!
!!! The whole SIAM oAuth process was deleted. The application is just mocking security !!!


# development start db
docker run --rm -d -p 5432:5432 -e POSTGRES_PASSWORD="1234" --name postgres postgres:alpine
