FROM golang:latest
LABEL authors="mpedrozoduran"

WORKDIR /orchestrator

RUN mkdir -p ./config

COPY ./config/config.yaml ./config/config.yaml
COPY ./deploy/app ./app

ENTRYPOINT /orchestrator/app