# Stage 1 - Builder: Import the golang container.
FROM golang:1.21.6-alpine as builder

ARG PORT=3000

# Install ssh client and git
RUN apk add --no-cache openssh-client git

# Download public key for github.com
RUN mkdir -p -m 0700 ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts

# Set the work directory.
WORKDIR /app

# Copy go mod and sum files.
COPY go.mod ./
COPY go.sum ./

# Install the dependencies.
RUN git config --global --add url."ssh://git@github.com/".insteadOf "https://github.com/"
# This command will have access to the forwarded agent (if one is
# available)
RUN go mod download

# Copy the source code into the container.
COPY ./ ./

# Build the source code
RUN CGO_ENABLED=0 go build -o ./out/application ./cmd/container/main.go


# Stage 2 - Runner.
FROM alpine:3.16.2
WORKDIR /app
COPY --from=builder /app/out/application application

EXPOSE $PORT
CMD [ "/app/application" ]