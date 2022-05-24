FROM golang:alpine3.15 as builder

ENV PATH=/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:/go/bin:/usr/local/go/bin

ENV CGO_ENABLED=0

RUN apk add --no-cache git
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3

COPY metricbeat/metricbeat /usr/share/beats/metricbeat
COPY metricbeat.yml /usr/share/beats/metricbeat.yml

FROM alpine:3.15

ENV PATH=/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/bin:/go/bin

WORKDIR /usr/share/beats

ENV ELASTICSEARCH_PASSWORD=changeme
ENV ELASTICSEARCH_USERNAME=elastic
ENV ELASTICSEARCH_HOST=elasticsearch

COPY --from=builder /go/bin/dlv /go/bin/dlv
COPY --from=builder /usr/share/beats/metricbeat /usr/share/beats/metricbeat
COPY --from=builder /usr/share/beats/metricbeat.yml /usr/share/beats/metricbeat.yml

CMD ["dlv", "--headless=true", "--listen=:56268", "--api-version=2", "--log", "exec", "./metricbeat", "--", "-c", "./metricbeat.yml", "-e"]