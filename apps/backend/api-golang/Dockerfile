#### BASE STAGE
FROM golang:1.25.3 AS base
WORKDIR /app

# Install moon binary
RUN curl -fsSL https://moonrepo.dev/install/moon.sh | bash
ENV PATH="/root/.moon/bin:/root/.proto/bin:$PATH"

#### BUILD STAGE
FROM base AS build

# Copy workspace configuration
COPY .moon .moon
COPY pnpm-workspace.yaml* ./
COPY package.json* ./

# Copy shared packages
COPY packages ./packages

# Copy api-golang project
COPY apps/backend/api-golang ./apps/backend/api-golang

# Install Go dependencies
WORKDIR /app/apps/backend/api-golang
RUN go mod download

# Build the application as a static binary for Alpine
RUN mkdir -p /app/bin && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/bin/api-golang ./cmd/api

#### RUNTIME STAGE
FROM golang:1.25.3-alpine AS runtime
WORKDIR /app

# Copy the binary
COPY --from=build /app/bin/api-golang /app/bin/api-golang

# Expose port
EXPOSE 8080

# Run the application
CMD ["/app/bin/api-golang"]

