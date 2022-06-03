FROM golang:1.17.2-alpine AS base
WORKDIR /src
COPY . .
WORKDIR /src/bin/state_machine_executor
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/app cmd/main.go

FROM alpine:latest
WORKDIR /bin
COPY --from=0 /src/bin/state_machine_executor/config/config.docker.yaml config/config.yaml
COPY --from=0 /bin/app app
ENTRYPOINT ./app
