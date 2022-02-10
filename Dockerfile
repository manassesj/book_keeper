# syntax=docker/dockerfile:1

FROM golang:1.16-buster as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-bookeeper

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /docker-bookeeper /docker-bookeeper

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-bookeeper"]