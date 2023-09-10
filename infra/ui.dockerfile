# syntax=docker/dockerfile:1

# Build stage
# ---
FROM node:18-alpine AS build
ARG API_ENDPOINT

ENV NODE_ENV production
ENV REACT_APP_API_ENDPOINT $API_ENDPOINT
WORKDIR /app

COPY ui/package.json .
COPY ui/package-lock.json .
RUN npm i --omit=dev

COPY ui .
RUN npm run build

# Run stage
# ---
FROM nginx:1.25.2-alpine

ENV NODE_ENV production

COPY --from=build /app/build /usr/share/nginx/html
COPY infra/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]