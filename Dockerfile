FROM golang

WORKDIR /data
COPY . /data/

RUN go build main.go

CMD ["/data/main"]