# Develop
FROM node:16.13.0 as dev-env

WORKDIR "/app"

EXPOSE 3000

# Build
FROM node:16.13.0 as build-env

WORKDIR /work
COPY . .
RUN yarn install
RUN yarn build

# Product
FROM node:16.13.0 as prod-env

WORKDIR /app
COPY --from=build-env /work/build /app/build
RUN npm install serve

EXPOSE 5000

ENTRYPOINT ["npx", "serve", "-s", "build"]
