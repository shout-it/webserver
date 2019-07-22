# Webserver
Web server for shoutit web blogger

# Dependencies
1. go - 1.10+
2. govendor(https://devcenter.heroku.com/articles/go-dependencies-via-govendor)

# How to run the application

1. Clone the application from `git@github.com:shout-it/webserver.git` to your local.
2. `bash scripts/hook-setup.sh` to set up pre-commit hooks.
3. Run `go get` or `dep ensure` to download the required packages.
4. `go run main.go` to start the application. It will start on `8080` port
5. `go test` to run unit test cases.
