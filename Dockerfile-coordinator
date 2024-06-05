FROM golang:1.22-alpine3.18

WORKDIR /workdir
COPY ./noble-load-tester /workdir/noble-load-tester
COPY ./tm-load-test/ /workdir/tm-load-test
COPY ./noblechain /workdir/noblechain
COPY ./start-coordinator.sh /workdir/

RUN chmod +x ./start-coordinator.sh 

WORKDIR /workdir/noble-load-tester

RUN apk add --no-cache --upgrade bash
RUN apk add --no-cache git

RUN go build -o ./build/noble-load-tester ./cmd/noble-load-tester/main.go

RUN chmod +x ./build/noble-load-tester

ENTRYPOINT ["../start-coordinator.sh"]

# CMD ["./build/noble-load-tester", "coordinator", "--expect-workers", "2", "--bind", "0.0.0.0:26670", "-c", "1", "-T", "1000", "-r", "10", "-s", "250", "--broadcast-tx-method", "async", "--endpoints", "ws://172.17.0.1:26657/websocket", "--client-factory", "noble"]

EXPOSE 26670
