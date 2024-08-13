##
## Build
##
FROM golang:1.21-alpine AS build

ARG USER
ARG TOKEN
ARG BUILD_PATH

ENV USER=$USER
ENV TOKEN=$TOKEN
ENV BUILD_PATH=$BUILD_PATH
ENV MIGRATION_PATH=$MIGRATION_PATH

WORKDIR /app

COPY . .

RUN apk --no-cache add g++ \
     make \
     bash \
     libstdc++ \
     curl \
     ca-certificates \
     protoc \
     protobuf \
     protobuf-dev \
     git \
     librdkafka \
     gcc \
     musl-dev

RUN export PATH="$PATH:/usr/.local/bin"
RUN export PATH="$PATH:$(go env GOPATH)/bin"

WORKDIR /app/src

RUN go mod download
RUN apk upgrade --no-cache

WORKDIR /app/src/$BUILD_PATH/cmd
RUN go build -ldflags '-extldflags "-static"' -tags musl -o /app/src/$BUILD_PATH/cmd/bin
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.12 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe
##
## Deploy
##
FROM alpine
#
ARG BUILD_PATH
ARG MIGRATION_PATH
ENV BUILD_PATH=$BUILD_PATH
ENV MIGRATION_PATH=$MIGRATION_PATH

COPY --from=build /app/src/sql/migrations/$MIGRATION_PATH /sql/migrations/$MIGRATION_PATH
COPY --from=build /app/src/$BUILD_PATH/cmd/bin /app/bin
COPY --from=build /app/src/$BUILD_PATH /src/$BUILD_PATH
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /app/docker/$MIGRATION_PATH/scripts /$MIGRATION_PATH/scripts
RUN /$MIGRATION_PATH/scripts/after-build.sh

EXPOSE 11001

CMD ["/app/bin"]
