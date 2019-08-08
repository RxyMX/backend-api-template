FROM golang:1.12.3 as builder

ARG GIT_USER
ARG GIT_TOKEN

WORKDIR /app

COPY . .

# add the token to be able to checkout private repo
RUN git config \
  --global \
  url."https://${GIT_USER}:${GIT_TOKEN}@github.com".insteadOf \
  "https://github.com"

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o main ./cmd/common-go-example/main.go

RUN rm -f ~/.gitconfig

# multi-stage build to reduce final image size
FROM alpine:3.9.3

# need this line to fix the error
# Post https://some/ssl/ednpoint: x509: certificate signed by unknown authority
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=builder /app/main main

ENTRYPOINT ["./main"]
