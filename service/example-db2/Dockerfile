FROM golang:1.19-buster as builder

RUN apt-get update -y && apt-get install libxml2 tzdata -y

WORKDIR /app
COPY ./service/example-db2 ./service/example-db2
COPY ./database ./database
COPY ./response ./response
COPY ./tracer ./tracer
COPY ./server ./server
COPY go.mod .
COPY go.sum .

# add ibm driver
RUN tar -xzf ./database/db2/linuxx64_odbc_cli.tar.gz -C /app

# ibm db2 configuration
ENV IBM_DB_HOME=/app/clidriver
ENV CGO_CFLAGS=-I$IBM_DB_HOME/include
ENV CGO_LDFLAGS=-L$IBM_DB_HOME/lib
ENV LD_LIBRARY_PATH=$IBM_DB_HOME/lib

RUN go build -o main ./service/example-db2

FROM ubuntu:20.04
WORKDIR /app

COPY --from=builder /app/main /app
COPY --from=builder /app/clidriver /app/clidriver

ENV IBM_DB_HOME=/app/clidriver
ENV CGO_CFLAGS=-I$IBM_DB_HOME/include
ENV CGO_LDFLAGS=-L$IBM_DB_HOME/lib
ENV LD_LIBRARY_PATH=$IBM_DB_HOME/lib

RUN apt update && apt -y install libxml2

CMD ["./main"]