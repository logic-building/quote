FROM golang:1.12.5 AS builder

WORKDIR /quote
COPY . .

RUN CGO_ENABLED=0 go install -v ./...

FROM alpine:3.8
COPY --from=builder /go/bin/quote /quote
COPY image /image
COPY image-motivational /image-motivational

RUN apk add --no-cache tzdata
ENV TZ=America/Denver

EXPOSE 1922
ENTRYPOINT ["/quote"]
