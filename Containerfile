FROM docker.io/golang:latest as build

WORKDIR /app
COPY Makefile .
COPY go.* .
RUN make deps

COPY . .
RUN make build

# cannot use scratch image as some bits are needed for the webserver
FROM docker.io/debian:stable-slim
COPY --from=build /app/bin/ready /app/bin/ready

ENTRYPOINT ["/app/bin/ready"]
