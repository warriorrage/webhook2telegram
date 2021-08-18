FROM golang:alpine

WORKDIR /data
COPY . .
RUN GO111MODULE=on go build -o webhook2telegram .

FROM alpine

WORKDIR /app
COPY --from=0 /data/webhook2telegram webhook2telegram
COPY --from=0 /data/views views
COPY --from=0 /data/version.txt version.txt
COPY --from=0 /data/docker/entrypoint.sh entrypoint.sh

ENV APP_TOKEN "02ef482f-2fa4-4aca-b76b-1cae2a11fa2e"
ENV APP_MODE "webhook"
ENV APP_METRICS "false"
ENV APP_RATE_LIMIT "190"

VOLUME /srv/data
EXPOSE 8080

ENTRYPOINT [ "/app/entrypoint.sh" ]
