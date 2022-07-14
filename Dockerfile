FROM golang:latest

COPY . /app

WORKDIR /app

RUN go get github.com/gorilla/mux
RUN go get github.com/leonelquinteros/gorand
RUN go get gorm.io/driver/postgres
RUN go get gorm.io/gorm

EXPOSE 3000

RUN go build



CMD ["./interview_restapi"]
