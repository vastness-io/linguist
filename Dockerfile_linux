From alpine:3.6

RUN apk add --no-cache  ca-certificates 

ADD bin/linux/amd64/linguist /linguist
EXPOSE 8082
ENTRYPOINT ["/linguist"]
