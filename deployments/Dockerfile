# Builder stage
FROM --platform=$BUILDPLATFORM golang:1.25.5 AS build
LABEL stage=intermediate
WORKDIR /app
COPY . .
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o bin/freeport ./cmd/freeport

# Runtime stage
FROM scratch
COPY --from=build /app/bin/freeport /bin/freeport
CMD ["freeport"]
