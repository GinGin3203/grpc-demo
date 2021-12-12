FROM golang:1.17-alpine as builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o users_service github.com/GinGin3203/grpc-demo/server
WORKDIR /dist
RUN cp /build/users_service .
FROM scratch
COPY --from=builder /dist/users_service /

ENTRYPOINT ["/users_service"]
