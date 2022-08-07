# what-the-food

Thai-Food classification mobile application

## Description

This is the `back-end` part, which is written in `Go-lang`.

## Installation

Use this command to init go mod

```bash
go mod init example.com/path
```

Install Gin-Gonic

```bash
go get -u github.com/gin-gonic/gin
```

Download the required dependencies

```bash
go mod download
go mod vertify
```

## Usage

_Run the app using go run command_

```bash
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
