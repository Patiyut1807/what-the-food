FROM golang:1.18

RUN go install github.com/beego/bee/v2@latest

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /what-the-food
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY src/app.py ./app.py
COPY src/class_food.txt ./class_food.txt
COPY src/model_food.pth ./model_food.pth
COPY src/requirements.txt ./requirements.txt

RUN apt-get update
RUN apt-get -y install python3-pip
RUN pip3 install -r requirements.txt

EXPOSE 8000
CMD ["bee", "run"]
