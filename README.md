<h1 align="center">MyGram API</h1>

## About The Project
MyGram API is a REST API built with the Gin server that allows you to post and comment on photos.\
This project implements the REST method and applies a clean architecture structure. So, it can be applied and scaled up easily at any time.

## Getting Started

### Prerequisites
- Go 1.8+
- PostgreeSQL 14.+

### Installation or Configure
- Clone the repo
```bash
# clone if you don't have the code base
$ git clone git@github.com:dhichii/mygram-api.git

# tidy up and run
$ go mod tidy
```
- Setup your env\
copy the `example.env`, rename to `app.env`, and define your own env
- run the app
```bash
$ go run main.go
```