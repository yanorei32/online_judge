# Develop
FROM node:14.17.2 as dev-env

WORKDIR "/app"

EXPOSE 3000

# Build
FROM node:14.17.2 as build-env

WORKDIR /work
COPY . .
RUN npm ci
RUN npm run build

# Product
FROM node:14.17.2 as prod-env

WORKDIR /app
COPY --from=build-env /work/build /app/build
RUN npm install serve

EXPOSE 5000

ENTRYPOINT ["npx", "serve", "-s", "build"]
