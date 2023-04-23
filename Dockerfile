# syntax=docker/dockerfile:1

FROM golang:1.19 as build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server

# TODO: Estudar distroless
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /app

# Copy server binary and .env file to /app
COPY --from=build /app/server /app/server
COPY --from=build /app/.env /app/.env

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT [ "/app/server" ]
