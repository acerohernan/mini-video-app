default: run

run:
	go run main.go

dev:
	CompileDaemon -build="go build main.go" -command="./main"