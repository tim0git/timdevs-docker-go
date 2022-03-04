FROM docker:20.10.12-alpine3.14

LABEL timdevs-docker-go="v1.0.0"

RUN apk -U upgrade

RUN apk add git

RUN apk add go

#RUN addgroup -S docker && adduser -S appuser -G docker
#
#USER appuser
