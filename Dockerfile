FROM golang:1.21.4 AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM scratch

COPY --from=builder /app/main /main

EXPOSE  6060

ENTRYPOINT ["/main"]
