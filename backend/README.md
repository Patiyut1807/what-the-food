# what-the-food

Thai-Food classification mobile application

## Description

This is the `back-end` part, which is written in `Go-lang`.

## Installation

In `src` floder

Use this command to init go mod

```go
go mod init example.com/path
```

Install Fiber V2

```go
go get -u github.com/gofiber/fiber/v2
```

Download the required dependencies

```go
go mod download
go mod vendor
go mod vertify
```

## Usage

_Run the app using go run command_

```go
go run main.go
```

_Build docker image_

```bash
docker build -t what-the-food-backend .
```

```bash
docker run -it -d --rm -p 8000:8000 -v $PWD/src:/go/src/what-the-food-backend what-the-food-backend
```

_Go to `http://server_pi:8000/test` and see the result_

```text
Hello
```