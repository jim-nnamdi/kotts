FROM golang:1.16-alpine

WORKDIR /app

ADD . /app

RUN CGO_ENABLED=1 

RUN GOOS= linux 

RUN go mod download 

RUN go mod tidy 

RUN go run -o main.go 

CMD [ "app", "main" ]