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

# Copy web-sdks-apps project files
COPY apps/frontend/web-sdks-apps/package.json ./apps/frontend/web-sdks-apps/
COPY apps/frontend/web-sdks-apps/pnpm-lock.yaml* ./apps/frontend/web-sdks-apps/

# Install dependencies
RUN pnpm install --frozen-lockfile

#### BUILD STAGE
FROM dependencies AS build

# Copy web-sdks-apps source
COPY apps/frontend/web-sdks-apps ./apps/frontend/web-sdks-apps

# Build the application
WORKDIR /app/apps/frontend/web-sdks-apps
RUN pnpm run build

#### RUNTIME STAGE
FROM node:22.20.0-alpine AS runtime
WORKDIR /app

# Copy package files and built application
COPY --from=build /app/apps/frontend/web-sdks-apps/package.json ./
COPY --from=build /app/apps/frontend/web-sdks-apps/build ./build
COPY --from=build /app/apps/frontend/web-sdks-apps/node_modules ./node_modules

# Expose port
EXPOSE 4001

# Set environment
ENV NODE_ENV=production
ENV PORT=4001

# Run the application (SvelteKit node adapter)
CMD ["node", "build"]

