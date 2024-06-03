FROM golang:1.20-alpine3.18

WORKDIR /workdir
COPY ../noble-load-tester /workdir/

RUN ls

COPY ../tm-load-test/ /workdir/
COPY ../noblechain /workdir/

WORKDIR /workdir/noble-load-tester

RUN apk add --no-cache git

RUN go build -o ./build/noble-load-tester ./cmd/noble-load-tester/main.go

RUN chmod +x ./build/noble-load-tester

EXPOSE 26670
