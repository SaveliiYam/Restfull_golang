build:
	docker-compose build todo-app

run: 
	docker-compose up todo-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:662011Egor@localhost:5436/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go