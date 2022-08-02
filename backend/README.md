# what-the-food

Thai-Food classification mobile application

## Description

This is the `back-end` part, which is written in `Go-lang`.

## Installation

Use this command to init go mod

```bash
go mod init example.com/path
```

Install Fiber V2

```bash
go get -u github.com/gofiber/fiber/v2
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
docker run -p 8000:8000 what-the-food-backend:latest
```
