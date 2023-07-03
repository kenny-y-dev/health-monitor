# syntax=docker/dockerfile:1

FROM golang:1.19 AS build-stage
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-health-monitor

FROM gcr.io/distroless/base-debian11 AS build-release-stage
WORKDIR /
COPY --from=build-stage /docker-health-monitor /docker-health-monitor
USER nonroot:nonroot
ENTRYPOINT ["/docker-health-monitor"]