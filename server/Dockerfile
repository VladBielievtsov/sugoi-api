FROM golang:1.22.4 as builder

RUN apt-get update && apt-get install -y libwebp-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /bin/sugoiapi cmd/api/main.go

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y libwebp-dev

ENV LD_LIBRARY_PATH=/usr/lib/x86_64-linux-gnu/

COPY --from=builder /bin/sugoiapi /bin/sugoiapi

COPY .env ./

EXPOSE 4000

CMD ["/bin/sugoiapi"]