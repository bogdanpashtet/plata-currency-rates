FROM golang:1.21.6-alpine3.18 as builder

RUN apk add --no-cache git

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/plata-currency-rates/main.go

FROM alpine:3.18
RUN apk update && apk add tzdata
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV DB_POSTGRES_USER=postgres
ENV DB_POSTGRES_PASSWORD=qwerty1234

EXPOSE 8080
EXPOSE 5432

COPY --from=builder /src/app .
COPY configs configs
COPY docs docs

CMD ["/app"]