# syntax=docker/dockerfile:1

FROM golang:1.21.5

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code, including subdirectories
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /itmo_delivery_bot

# Run
CMD ["/itmo_delivery_bot"]
