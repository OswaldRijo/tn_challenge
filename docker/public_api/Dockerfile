FROM node:22.6.0-alpine3.20 AS base

# Install dependencies only when needed
FROM base AS deps

ARG USER
ARG BUILD_PATH

ENV USER=$USER
ENV BUILD_PATH=$BUILD_PATH

# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY src/node/$BUILD_PATH/.npmrc ./
# Install dependencies based on the preferred package manager
COPY src/node/$BUILD_PATH/package.json ./
COPY src/node/$BUILD_PATH/pnpm-lock.yaml ./
RUN npm -g add pnpm && pnpm i --frozen-lockfile;

# Rebuild the source code only when needed
FROM base AS builder
ARG BUILD_PATH
ENV BUILD_PATH=$BUILD_PATH

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY src/node/$BUILD_PATH ./

# Next.js collects completely anonymous telemetry.ts data about general usage.
# Learn more here: https://nextjs.org/telemetry
# Uncomment the following line in case you want to disable telemetry.ts during the build.
# ENV NEXT_TELEMETRY_DISABLED 1
RUN npm -g add pnpm
RUN pnpm run build

# If using npm comment out above and use below instead
# RUN npm run build

# Production image, copy all the files and run next
FROM base AS runner
WORKDIR /app

ENV NODE_ENV production
# Uncomment the following line in case you want to disable telemetry.ts during runtime.
# ENV NEXT_TELEMETRY_DISABLED 1

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 nestjs


# Automatically leverage output traces to reduce image size
# https://nextjs.org/docs/advanced-features/output-file-tracing
COPY --from=deps /app/node_modules ./node_modules
COPY --from=builder --chown=nestjs:nodejs /app/dist /app/dist

USER nestjs

EXPOSE 8080

ENV PORT 8080

CMD ["node", "/app/dist/main.js"]
