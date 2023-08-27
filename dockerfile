FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/adnicolas/golang-hexagonal
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/gohex-api cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/gohex-api /go/bin/gohex-api
ENTRYPOINT ["/go/bin/gohex-api"]