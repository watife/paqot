# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Fakorede Boluwatife"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container
COPY . .

ENV GO111MODULE=on

RUN go get ./...

# Build the Go api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .
#COPY --from=builder /app ./templates
COPY --from=builder /app/pkg ./pkg
#COPY --from=builder /api/migrations ./migrations
# COPY --from=builder /api/tls ./tls
#COPY --from=builder /app/.env.docker ./.env


# Expose port to the outside world
#EXPOSE 4000

#Command to run the executable
CMD [ "./main" ]