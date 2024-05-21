#!/bin/sh

cd ./address-book-backend
GOOS=linux GOARCH=amd64 go build -o app

cd ../address-book-frontend
pnpm run build:ci

cd ..
docker build -t addressbook .  
