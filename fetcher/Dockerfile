FROM golang:1.12

RUN mkdir -p /build
WORKDIR /build/
COPY . .
RUN GOOS=linux go build -mod vendor -o bin/fetcher github.com/Ezetowers/tennis-statistics/fetcher

FROM golang:1.12
COPY --from=0 /build/bin/fetcher /usr/local/sbin/fetcher
ENTRYPOINT ["/bin/sh"]