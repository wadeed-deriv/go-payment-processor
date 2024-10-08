

# Use the official Golang image as a build environment
FROM golang:1.23 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/gopaymentprocessor/main.go

FROM golang:1.20

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
