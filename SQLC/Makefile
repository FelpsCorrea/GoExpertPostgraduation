createmigration:
	migrate create -ext sql -dir sql/migrations -seq init

migrate:
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3307)/courses" -verbose up

migratedown: 
	migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3307)/courses" -verbose down

# Excluir a possibilidade de rodar o arquivo migrate quando rodar o make
.PHONY: migrate migratedown createmigration