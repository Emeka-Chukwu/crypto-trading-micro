FROM golang:1.21.4-alpine

WORKDIR /app

COPY go.mod ./
# COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /notifications-service

EXPOSE 8080

CMD [ "/notifications-service" ]
