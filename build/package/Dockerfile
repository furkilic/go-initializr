FROM alpine:3.11.5

ARG EXEC_DIR="/bin/go-initializr"
ARG ASSETS="/assets/templates"

COPY $EXEC_DIR /data/app/go-initializr
COPY $ASSETS /data/app/assets/templates

WORKDIR "/data/app/"

CMD "/data/app/go-initializr"