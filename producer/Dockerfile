FROM golang:1.12

RUN mkdir -p /build
WORKDIR /build/
COPY . .
RUN GOOS=linux go build -mod vendor -o bin/producer github.com/seguijoaquin/tennis-statistics/producer

FROM golang:1.12
COPY --from=0 /build/bin/producer /usr/local/sbin/producer
ENTRYPOINT ["/bin/sh"]