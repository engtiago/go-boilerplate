run:
	CompileDaemon -command="./go-boilerplate"

build:
	go build -o dist/api.exe main.go

migratedb:
	go run migrate/migrate.go

createdocker:
	docker image rm go-server -f && docker build . -t go-server

rundocker:
	docker rm goserver -f && docker run --name goserver -p 3000:3000 go-server

createfulldocker:
	 make createdocker && make rundocker

generateimage:
	docker save -o go-server.tar go-server:latest

gitlabserver:
	docker build . -t registry.gitlab.com/engtiago/go-boilerplate/go-server && docker push registry.gitlab.com/engtiago/go-boilerplate/go-server

rungitlabdocker:
	docker compose up -d