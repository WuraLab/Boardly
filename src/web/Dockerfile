FROM node:alpine AS builder
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY . /app
RUN yarn install --frozen-lockfile
RUN yarn build


FROM node:alpine AS runner
WORKDIR /app

ENV NODE_ENV production

COPY --from=builder /app/src/web/next.config.js ./
COPY --from=builder /app/src/web/public ./public
COPY --from=builder /app/src/web/.next ./.next

COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001
RUN chown -R nextjs:nodejs /app/.next
USER nextjs

EXPOSE 3000

CMD ["yarn", "start"]
