ARG GOLANG_VERSION=1.20.4

### GOLANG BUILD
FROM golang:${GOLANG_VERSION}-alpine as build

RUN apk --no-cache add ca-certificates
RUN apk --no-cache add \
    git \
    git-lfs \
    ;

WORKDIR /go/src/web

COPY . .

ENV CGO_ENABLED 0

RUN go mod download && go mod verify && go install -v ./... ;

### GOLANG PROD
FROM alpine:latest AS prod

WORKDIR /go/src/web

RUN apk add --no-cache tzdata;

COPY --from=build /go/src/web/app /go/src/web/app
COPY --from=build /go/bin/app /go/bin/app
EXPOSE 8000
CMD ["/go/bin/app"]
