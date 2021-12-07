FROM golang:1.17-alpine as builder
WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 go build -o server

FROM scratch
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --chown=0:0 --from=builder /build/* /
USER 65534
EXPOSE 3000:3000
ENTRYPOINT [ "/server" ]
