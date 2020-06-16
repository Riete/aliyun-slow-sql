FROM riet/golang:1.13.10 as backend
COPY . .
RUN unset GOPATH && go build -mod=vendor

FROM riet/centos:7.4.1708-cnzone
COPY --from=backend /go/aliyun-slow-sql /opt/aliyun-slow-sql
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]