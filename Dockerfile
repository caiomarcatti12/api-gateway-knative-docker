FROM golang:1.19.0 as dev
WORKDIR /app

FROM dev as build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/api-gateway-knative-docker

FROM alpine:3.16.2 as prod
WORKDIR /app

COPY --from=build /app/api-gateway-knative-docker /app/api-gateway-knative-docker

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
RUN chown -R appuser:appgroup /app && \
    chmod +x /app/api-gateway-knative-docker
USER appuser

RUN apk add bash
RUN ls -lha
#CMD ["/app/api-gateway-knative-docker"]
CMD ["/bin/bash"]
HEALTHCHECK --interval=10s --timeout=3s CMD curl -f http://localhost:8080/ || exit 1