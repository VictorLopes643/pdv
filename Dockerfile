FROM golang:latest as builder

WORKDIR /app

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o api .

FROM scratch
COPY --from=builder /app/api .
CMD ["./api"]