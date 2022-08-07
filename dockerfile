FROM golang:1.16-alpine 

WORKDIR /app 

ADD . /app/   

RUN CGO_ENABLED=1 \ 
    GOOS=linux \
    go mod download 

RUN go mod download golang.org/x/net

RUN go get github.com/githubnemo/CompileDaemon

COPY entrypoint.sh /app/

RUN go build -o main . 

CMD [ "sh", "entrypoint.sh" ]