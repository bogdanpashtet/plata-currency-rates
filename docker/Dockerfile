FROM golang:1.21.4-alpine3.18 as builder
ARG VERSION

RUN apk update && apk add gcc g++ make bash git \
libevent-dev zlib-dev libressl-dev \
zstd-dev zstd-static

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/plata-currency-rates/main.go

FROM alpine:3.18
RUN apk update && apk add tzdata
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV GOAPP_ENVIROMENMENT=local

EXPOSE 8080

COPY --from=builder /src/app .
COPY configs configs

CMD ["/app"]