#window
migrate -source file:api\data\migrations -database postgres://postgres:admin@localhost:5432/postgres?sslmode=disable up 2

#mac
migrate -path api/data/migrations -database "postgresql://postgres:admin@localhost:5432/postgres?sslmode=disable" -verbose up 2


#docker down 2
migrate -source file:db\migration -database postgres://postgres:admin@localhost:5433/postgres?sslmode=disable down 2

