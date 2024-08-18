FROM node:22.6.0-alpine3.20 AS base
# Install dependencies only when needed
FROM base AS deps
RUN apk add --no-cache libc6-compat
WORKDIR /app

COPY src/node/web_app/.npmrc ./
# Install dependencies based on the preferred package manager
COPY src/node/web_app/package.json pnpm-lock.yaml* ./
RUN npm -g add pnpm && pnpm i  --frozen-lockfile;


# Rebuild the source code only when needed
FROM base AS builder
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
COPY src/node/web_app .

RUN npm -g add pnpm
RUN pnpm build --no-lint

# If using npm comment out above and use below instead
# RUN npm run build

# Production image, copy all the files and run next
FROM node:22.6.0-alpine3.20
WORKDIR /app

ENV NODE_ENV production
# Uncomment the following line in case you want to disable telemetry.ts during runtime.
# ENV NEXT_TELEMETRY_DISABLED 1

RUN addgroup --system --gid 1001 nodejs
RUN adduser --system --uid 1001 reactjs


COPY --from=builder /app/build /app/build
COPY --from=builder /app/server /app/server

RUN cd server && npm -g add pnpm  && pnpm i

# Expose port 80 for the Nginx server
EXPOSE 80

# Start Nginx when the container runs
CMD ["node", "/app/server/index.js"]
