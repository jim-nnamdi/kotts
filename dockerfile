FROM golang:1.16-alpine 

WORKDIR /app 

ADD . /app/   

RUN CGO_ENABLED=1 GOOS=linux go mod download

RUN ["apt-get", "update"]

RUN ["apt-get", "-y", "install", "vim"]

RUN go build -o main . 

CMD [ "/app/main" ]