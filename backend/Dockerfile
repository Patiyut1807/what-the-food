FROM golang:1.18

RUN apt-get update
RUN apt-get -y install python3-pip

RUN mkdir /wtf

WORKDIR /wtf

COPY model/app.py .
COPY model/class_food.txt .
COPY model/model_food.pth .
COPY model/requirements.txt .

COPY src/main.go .
COPY src/go.sum .
COPY src/go.mod .

RUN pip3 install -r requirements.txt

EXPOSE 5000
CMD ["go","run","main.go"]
