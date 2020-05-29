FROM golang:1.14 as base
ENV HTTP_PORT 8080
EXPOSE $HTTP_PORT
WORKDIR /go/src/github.com/diegodesousas/apistarter

FROM base as developer
RUN go get github.com/liudng/dogo

FROM base AS compiler
COPY . ./
RUN go mod vendor
RUN CGO_ENABLED=0 GOARCH=amd64 go build -o /apistarter/bin /go/src/github.com/diegodesousas/apistarter/cmd

FROM alpine:3.10 AS release
RUN apk add --update --no-cache ca-certificates tzdata && \
    cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && \
    rm -rf /var/cache/apk/* /tmp/* /var/tmp/* && \
    date

COPY --from=compiler /apistarter/bin /apistarter/bin

RUN addgroup -g 1000 -S hu && \
    adduser -u 1000 -S hu -G hu
USER hu

CMD /apistarter/bin