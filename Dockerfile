FROM golang:alpine as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build -o pns cmd/main.go

FROM alpine
RUN adduser -S -D -H -h /app service
USER service
COPY --from=builder /build/pns /app/
WORKDIR /app
ENTRYPOINT ["./pns"]
CMD ["--help"]