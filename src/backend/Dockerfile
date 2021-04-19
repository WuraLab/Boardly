FROM golang:1.15 AS builder
COPY . /go/src/github.com/wuraLab/boardly
WORKDIR /go/src/github.com/wuraLab/boardly
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app src/backend/cmd/boardly.go

FROM scratch
COPY --from=builder /app .
EXPOSE 8000
CMD ["/app"]

