# Mini-Payment-Gateway
Golang Practice REST API

## Order-MG
An Mini Payment Gateway

## Description
This is my REST API Golang practice with a basic transaction behaviors and CRUD on Transaction. 

## Technologies
- Golang 1.19
- Gochi 
- GORM 
- PostgreSQL
- SonyFlake
- Logrus
- Mockery
- Testify
- Docker

## Executing Program
### to migrate the db:  
migrate -source file:api\data\migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up 2

### to run the program: 
go run .api\cmd\served\main.go

### to migrate and run program by docker: 
docker-compose up db app migrate_up

## Postman test
https://www.getpostman.com/collections/9076164695430ac3cd3a

####
