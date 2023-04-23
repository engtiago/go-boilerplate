run:
	CompileDaemon -command="./go-boilerplate"

build:
	go build -o dist/api.exe main.go

migratedb:
	go run migrate/migrate.go

createdocker:
	docker build . -t go-server

rundocker:
	docker rm goserver -f && docker run --name goserver -p 3000:3000 go-server

createfulldocker:
	docker image rm go-server -f && make createdocker && make rundocker