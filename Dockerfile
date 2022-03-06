FROM golang:latest as builder

RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app
RUN go build -o server cmd/blockhain/server/main.go
RUN go build -o client cmd/blockhain/client/main.go

FROM debian:buster-slim

RUN mkdir -p /data
ENV BLOCKCHAIN_FILEPATH /data/quotes.txt
COPY internal/repository/in_memory/quotes.txt ${BLOCKCHAIN_FILEPATH}
COPY --from=builder /go/src/app/server /usr/local/bin
COPY --from=builder /go/src/app/client /usr/local/bin
RUN chmod +x /usr/local/bin/server /usr/local/bin/client

EXPOSE 8080

