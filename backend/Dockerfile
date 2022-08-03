FROM golang:1.18

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_PATH /go/src/what-the-food-backend
RUN mkdir -p "$APP_PATH"

WORKDIR "$APP_PATH"
EXPOSE 8000
CMD ["bee", "run"]