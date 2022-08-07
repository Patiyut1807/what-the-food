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

Install Gin-Gonic

```go
go get -u github.com/gin-gonic/gin
```

Download the required dependencies

```go
go mod download
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
docker run -it -d --rm -p 5000:5000 what-the-food-backend
```

_Go to `http://server_pi:5000/test` and see the result_

```text
{ text: "Golang" }
```
