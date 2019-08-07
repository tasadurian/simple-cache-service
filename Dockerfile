FROM golang:1.12-alpine3.9 as BUILD

RUN apk add --no-cache git gcc musl-dev

WORKDIR /cache
ENV GOFLAGS=-mod=vendor
COPY ./ ./

WORKDIR /cache
RUN go build .

FROM alpine

COPY --from=BUILD /cache/simple-cache-service /cache/simple-cache-service
EXPOSE 7771
CMD ["/cache/simple-cache-service"]