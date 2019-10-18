FROM golang:1.12

RUN mkdir -p /build
WORKDIR /build/
COPY . .
RUN GOOS=linux go build -mod vendor -o bin/consumer github.com/seguijoaquin/tennis-statistics/consumer

FROM golang:1.12
COPY --from=0 /build/bin/consumer /usr/local/sbin/consumer
ENTRYPOINT ["/bin/sh"]