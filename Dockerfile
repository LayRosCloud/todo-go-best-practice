FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV APP_ENV=production

RUN apk update --no-cache

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o main ./cmd/api/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/migrations/ /app/migrations/

EXPOSE 8080

CMD [ "./main" ]