FROM alpine:latest

WORKDIR /root/

COPY bin/client .

ENTRYPOINT ["./client"]



