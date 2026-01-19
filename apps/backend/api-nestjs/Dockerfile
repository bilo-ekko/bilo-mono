#### BASE STAGE
FROM node:22.20.0 AS base
WORKDIR /app

# Enable corepack for pnpm
RUN corepack enable

#### DEPENDENCIES STAGE
FROM base AS dependencies

# Copy workspace configuration
COPY .moon .moon
COPY pnpm-workspace.yaml ./
COPY pnpm-lock.yaml ./
COPY package.json* ./

# Copy shared packages
COPY packages ./packages

# Copy api-nestjs project files
COPY apps/backend/api-nestjs/package.json ./apps/backend/api-nestjs/
COPY apps/backend/api-nestjs/pnpm-lock.yaml* ./apps/backend/api-nestjs/

# Install dependencies
RUN pnpm install --frozen-lockfile

#### BUILD STAGE
FROM dependencies AS build

# Copy api-nestjs source
COPY apps/backend/api-nestjs ./apps/backend/api-nestjs

# Build the application
WORKDIR /app/apps/backend/api-nestjs
RUN pnpm run build

#### RUNTIME STAGE
FROM node:22.20.0-alpine AS runtime
WORKDIR /app

# Copy package files and built application
COPY --from=build /app/apps/backend/api-nestjs/package.json ./
COPY --from=build /app/apps/backend/api-nestjs/dist ./dist
COPY --from=build /app/apps/backend/api-nestjs/node_modules ./node_modules

# Expose port
EXPOSE 3000

# Run the application
CMD ["node", "dist/main.js"]

