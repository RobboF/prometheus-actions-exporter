# syntax=docker/dockerfile:1

FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ${PWD} ./
RUN CGO_ENABLED=0 GOOS=linux && go build -o prometheus-actions-exporter ./cmd/main/main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/prometheus-actions-exporter /prometheus-actions-exporter

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/prometheus-actions-exporter"]
