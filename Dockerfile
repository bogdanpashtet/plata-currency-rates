FROM golang:1.21.6-alpine3.18 as builder

RUN apk add --no-cache git

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/plata-currency-rates/main.go

FROM alpine:3.18
RUN apk update && apk add tzdata
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 8080

COPY --from=builder /src/app .
COPY configs configs
COPY docs docs
COPY backup.sql backup.sql
COPY wait-for-postgres.sh wait-for-postgres.sh

# install psql
RUN apk add postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

CMD ["/app"]