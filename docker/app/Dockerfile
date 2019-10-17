# Build environment
# -----------------
FROM golang:1.13-alpine as build-env
WORKDIR /snippetbox

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

RUN mkdir tls && cd tls
RUN go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
RUN cd ..

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/web* \
  && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate

# Deployment environment
# ----------------------
FROM alpine
RUN apk update && apk add --no-cache bash mysql-client

COPY --from=build-env /snippetbox/bin/app /snippetbox/
COPY --from=build-env /snippetbox/bin/migrate /snippetbox/
COPY --from=build-env /snippetbox/migrations /snippetbox/migrations
COPY --from=build-env /snippetbox/tls /snippetbox/tls
COPY --from=build-env /snippetbox/ui /snippetbox/ui

COPY --from=build-env /snippetbox/docker/app/bin /usr/local/bin/snippetbox/
RUN chmod +x /usr/local/bin/snippetbox/*

WORKDIR /snippetbox

EXPOSE 8080
CMD ["./app"]