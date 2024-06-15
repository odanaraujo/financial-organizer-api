FROM golang:1.22.4 as BUILDER

WORKDIR /app
COPY . .
COPY cmd/api/main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o financial

FROM golang:1.22-alpine

COPY --from=BUILDER /app/financial ./financial

CMD ["./financial"]