# Start with a base image that has Go 1.19 installed
FROM golang:1.19-alpine as build

# Set the working directory to /app
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o myapp .

FROM alpine:latest AS final

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/myapp .

EXPOSE 8080

CMD ["./myapp"]
