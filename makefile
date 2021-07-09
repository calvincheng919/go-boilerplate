install: 
	go mod tidy

start:
	go run src/main.go

dev:
	ENV="local" go run src/main.go

testserver:
	ENV="test" go run src/main.go

test: 
	BASE_URL=http://localhost:8080 go test ./src

build:
	rm main; goos=linux go build src/main.go

zip:
	rm main main.zip; goos=linux go build src/main.go && zip main main

golib:
	GOPRIVATE=github.com/lemonbrew/golib go get github.com/lemonbrew/golib

doc:
	apidoc -o docs

init:
	rm go.mod && go mod init ${module} &&	sed -i 's/my-service/${module}/' src/*.go && rm -rf .git && git init .


