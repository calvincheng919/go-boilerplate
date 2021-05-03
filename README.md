## Start
- `$ git clone git@github.com:lemonbrew/go-service-boilerplate.git`
- `$ cd go-service-boilerplate`  
- `$ go mod tidy` (similar to npm install)
- `$ go run src/main.go`

## Run tests (after you've started server)
- `$ export BASE_URL=http://localhost:8080`
- `$ go test ./src`

## Update Documentation
- `$ npm install -g apidoc`
- `$ apidoc -o docs`

## Converting from boilerplate to specific service
- Replace boilderplate module name with actual module name 
  - `$ rm go.mod`
  - `$ go mod init my-actual-service`
  - Change imports to reflect new name in `main.go` and `main_test.go`
- Replace boilerplate git repo with actual git repo
  - `$ rm -rf .git`
  - `$ git init .`
