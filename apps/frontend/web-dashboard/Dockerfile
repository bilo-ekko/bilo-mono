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

# Copy web-dashboard project files
COPY apps/frontend/web-dashboard/package.json ./apps/frontend/web-dashboard/
COPY apps/frontend/web-dashboard/pnpm-lock.yaml* ./apps/frontend/web-dashboard/

# Install dependencies
RUN pnpm install --frozen-lockfile

#### BUILD STAGE
FROM dependencies AS build

# Copy web-dashboard source
COPY apps/frontend/web-dashboard ./apps/frontend/web-dashboard

# Build the application
WORKDIR /app/apps/frontend/web-dashboard
RUN pnpm run build

#### RUNTIME STAGE
FROM node:22.20.0-alpine AS runtime
WORKDIR /app

# Enable corepack for pnpm
RUN corepack enable

# Copy package files
COPY --from=build /app/apps/frontend/web-dashboard/package.json ./
COPY --from=build /app/apps/frontend/web-dashboard/next.config.ts ./
COPY --from=build /app/apps/frontend/web-dashboard/pnpm-lock.yaml* ./

# Copy public assets and built application
COPY --from=build /app/apps/frontend/web-dashboard/public ./public
COPY --from=build /app/apps/frontend/web-dashboard/.next ./.next
COPY --from=build /app/apps/frontend/web-dashboard/node_modules ./node_modules

# Expose port
EXPOSE 3000

# Set environment
ENV NODE_ENV=production
ENV PORT=3000

# Run the application
CMD ["pnpm", "start"]

