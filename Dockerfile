FROM golang:1.17-alpine as builder
WORKDIR /build
ADD . .
RUN CGO_ENABLED=0 go build -o server

FROM scratch
COPY --chown=0:0 --from=builder /build/server /server
COPY --chown=0:0 --from=builder /build/index.html /index.html
USER 65534
EXPOSE 3000:3000
ENTRYPOINT [ "/server" ]
