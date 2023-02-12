FROM alpine:latest

WORKDIR /root/

COPY bin/client .

ENTRYPOINT ["./client"]

CMD ["--svc-url=value"]

