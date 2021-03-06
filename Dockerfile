FROM golang:1.15 AS builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app .

FROM scratch
COPY --from=builder /app .
EXPOSE 8000
CMD ["/app"]

